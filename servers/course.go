package servers

import (
	"context"

	"algohub.dev/backend/model"
	pb "algohub.dev/backend/proto"
	"algohub.dev/backend/structs"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"
)

type CourseServer struct {
	pb.CourseServiceServer
	DB  *gorm.DB
	Env *structs.Env
}

// GetCourses returns all courses
func (s *CourseServer) GetCourses(ctx context.Context, in *emptypb.Empty) (*pb.GetCoursesResponse, error) {
	courses := []*model.Course{}
	s.DB.Model(&model.Course{}).Find(&courses)

	// Convert courses to protobuf
	coursesProto := make([]*pb.GetCoursesResponse_Course, len(courses))
	for i, course := range courses {
		coursesProto[i] = &pb.GetCoursesResponse_Course{
			Id:          course.ID.String(),
			Name:        course.Name,
			Description: course.Description,
		}
	}

	return &pb.GetCoursesResponse{Courses: coursesProto}, nil
}

// GetCourse returns a course by id
func (s *CourseServer) GetCourse(ctx context.Context, in *pb.GetCourseRequest) (*pb.GetCourseResponse, error) {
	// Parse the UUID
	id, err := uuid.Parse(in.Id)
	if err != nil {
		errorText := "Invalid UUID"
		return &pb.GetCourseResponse{Success: false, Reason: &errorText}, nil
	}

	course := &model.Course{ID: id, DeletedAt: gorm.DeletedAt{}}
	result := s.DB.Preload("Chapters").Preload("Chapters.Sections").First(&course)

	// Check if course exists
	if result.Error != nil {
		errorText := "Course not found"
		return &pb.GetCourseResponse{Success: false, Reason: &errorText}, nil
	}

	// Conver chapters to protobuf
	chaptersProto := make([]*pb.Course_Chapter, len(course.Chapters))
	for i, chapter := range course.Chapters {
		sectionsProto := make([]*pb.Course_Chapter_Section, len(chapter.Sections))
		for j, section := range chapter.Sections {
			sectionsProto[j] = &pb.Course_Chapter_Section{
				Id:      section.ID.String(),
				Name:    section.Name,
				Order:   int32(section.Order),
				Title:   section.Title,
				Content: section.Content,
			}
		}
		chaptersProto[i] = &pb.Course_Chapter{
			Id:       chapter.ID.String(),
			Name:     chapter.Name,
			Order:    int32(chapter.Order),
			Sections: sectionsProto,
		}
	}

	// Convert course to protobuf
	courseProto := &pb.Course{
		Id:          course.ID.String(),
		Name:        course.Name,
		Description: course.Description,
		Chapters:    chaptersProto,
	}

	return &pb.GetCourseResponse{Course: courseProto, Success: true}, nil
}
