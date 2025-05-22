package anonymousstruct

import "fmt"

type Student struct {
	PersonalDetails struct { //anonymous struct
		Name          string
		Enrollment_Id int
	}
	GPA float64
}

/* anonymous Fields
Rules
1. Uniqueness Required
*/
type NewStudent struct {
	int
	string
	Address string //2.combining name and anonymous fiels
}

func StudentData() {
	student := Student{ //accessing anonymous struct
		PersonalDetails: struct {
			Name          string
			Enrollment_Id int
		}{"AIsh", 87},
	}
	fmt.Println(student)

	NewStudent := NewStudent{12, "NewStudent", "Bangalore"}
	fmt.Println(NewStudent)
}
