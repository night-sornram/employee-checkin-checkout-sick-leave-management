package repository

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockAttendanceRepo struct {
	GetAllFunc     func() ([]Attendance, error)
	GetByIDFunc    func(id int) (Attendance, error)
	CreateFunc     func(attendance Attendance) (Attendance, error)
	UpdateFunc     func(id int, attendance Attendance) (Attendance, error)
	DeleteFunc     func(id int) error
	CheckInFunc    func(eid string) (Attendance, error)
	CheckOutFunc   func(id int) (Attendance, error)
	GetAllMeFunc   func(eid string) ([]Attendance, error)
	CheckTodayFunc func(eid string) (Attendance, error)
}

func (m *mockAttendanceRepo) GetAll() ([]Attendance, error) {
	if m.GetAllFunc != nil {
		return m.GetAllFunc()
	}
	return nil, errors.New("not implemented")
}

func (m *mockAttendanceRepo) GetByID(id int) (Attendance, error) {
	if m.GetByIDFunc != nil {
		return m.GetByIDFunc(id)
	}
	return Attendance{}, errors.New("not implemented")
}

func (m *mockAttendanceRepo) Create(attendance Attendance) (Attendance, error) {
	if m.CreateFunc != nil {
		return m.CreateFunc(attendance)
	}
	return Attendance{}, errors.New("not implemented")
}

func (m *mockAttendanceRepo) Update(id int, attendance Attendance) (Attendance, error) {
	if m.UpdateFunc != nil {
		return m.UpdateFunc(id, attendance)
	}
	return Attendance{}, errors.New("not implemented")
}

func (m *mockAttendanceRepo) Delete(id int) error {
	if m.DeleteFunc != nil {
		return m.DeleteFunc(id)
	}
	return errors.New("not implemented")
}

func (m *mockAttendanceRepo) CheckIn(eid string) (Attendance, error) {
	if m.CheckInFunc != nil {
		return m.CheckInFunc(eid)
	}
	return Attendance{}, errors.New("not implemented")
}

func (m *mockAttendanceRepo) CheckOut(id int) (Attendance, error) {
	if m.CheckOutFunc != nil {
		return m.CheckOutFunc(id)
	}
	return Attendance{}, errors.New("not implemented")
}

func (m *mockAttendanceRepo) GetAllMe(eid string) ([]Attendance, error) {
	if m.GetAllMeFunc != nil {
		return m.GetAllMeFunc(eid)
	}
	return []Attendance{}, errors.New("not implemented")
}
func (m *mockAttendanceRepo) CheckToday(eid string) (Attendance, error) {
	if m.CheckTodayFunc != nil {
		return m.CheckTodayFunc(eid)
	}
	return Attendance{}, errors.New("not implemented")
}

func TestGetAll(t *testing.T) {
	t.Run("Valid-GetAll", func(t *testing.T) {
		mockRepo := &mockAttendanceRepo{
			GetAllFunc: func() ([]Attendance, error) {
				return []Attendance{}, nil
			},
		}

		service := NewAttendanceService(mockRepo)
		_, err := service.GetAttendances()
		assert.NoError(t, err)
	})
	t.Run("Invalid-GetAll", func(t *testing.T) {
		mockRepo := &mockAttendanceRepo{
			GetAllFunc: func() ([]Attendance, error) {
				return []Attendance{}, errors.New("invalid")
			},
		}

		service := NewAttendanceService(mockRepo)
		_, err := service.GetAttendances()
		assert.Error(t, err)
		assert.Equal(t, "invalid", err.Error())
	})
}

func TestGetByID(t *testing.T) {
	t.Run("Valid-GetByID", func(t *testing.T) {
		mockRepo := &mockAttendanceRepo{
			GetByIDFunc: func(id int) (Attendance, error) {
				return Attendance{}, nil
			},
		}

		service := NewAttendanceService(mockRepo)
		_, err := service.GetAttendance(1)
		assert.NoError(t, err)
	})
	t.Run("Invalid-GetByID", func(t *testing.T) {
		mockRepo := &mockAttendanceRepo{
			GetByIDFunc: func(id int) (Attendance, error) {
				return Attendance{}, errors.New("invalid")
			},
		}

		service := NewAttendanceService(mockRepo)
		_, err := service.GetAttendance(1)
		assert.Error(t, err)
		assert.Equal(t, "invalid", err.Error())
	})
}

func TestCreate(t *testing.T) {
	t.Run("Valid-Create", func(t *testing.T) {
		mockRepo := &mockAttendanceRepo{
			CreateFunc: func(attendance Attendance) (Attendance, error) {
				return Attendance{}, nil
			},
		}
		service := NewAttendanceService(mockRepo)
		_, err := service.CreateAttendance(Attendance{})
		assert.NoError(t, err)
	})
	t.Run("Valid-Create", func(t *testing.T) {
		mockRepo := &mockAttendanceRepo{
			CreateFunc: func(attendance Attendance) (Attendance, error) {
				return Attendance{}, errors.New("invalid")
			},
		}
		service := NewAttendanceService(mockRepo)
		_, err := service.CreateAttendance(Attendance{})
		assert.Error(t, err)
		assert.Equal(t, "invalid", err.Error())
	})
}
func TestUpdate(t *testing.T) {
	t.Run("Valid-Update", func(t *testing.T) {
		mockRepo := &mockAttendanceRepo{
			UpdateFunc: func(id int, attendance Attendance) (Attendance, error) {
				return Attendance{}, nil
			},
		}
		service := NewAttendanceService(mockRepo)
		_, err := service.UpdateAttendance(1, Attendance{})
		assert.NoError(t, err)
	})
	t.Run("Invalid-Update", func(t *testing.T) {
		mockRepo := &mockAttendanceRepo{
			UpdateFunc: func(id int, attendance Attendance) (Attendance, error) {
				return Attendance{}, errors.New("invalid")
			},
		}
		service := NewAttendanceService(mockRepo)
		_, err := service.UpdateAttendance(1, Attendance{})
		assert.Error(t, err)
		assert.Equal(t, "invalid", err.Error())
	})
}

func TestDelete(t *testing.T) {
	t.Run("Valid-Delete", func(t *testing.T) {
		mockRepo := &mockAttendanceRepo{
			DeleteFunc: func(id int) error {
				return nil
			},
		}
		service := NewAttendanceService(mockRepo)
		err := service.DeleteAttendance(1)
		assert.NoError(t, err)
	})
	t.Run("Invalid-Delete", func(t *testing.T) {
		mockRepo := &mockAttendanceRepo{
			DeleteFunc: func(id int) error {
				return errors.New("invalid")
			},
		}
		service := NewAttendanceService(mockRepo)
		err := service.DeleteAttendance(1)
		assert.Error(t, err)
		assert.Equal(t, "invalid", err.Error())
	})
}

func TestCheckIn(t *testing.T) {
	t.Run("Valid-CheckIn", func(t *testing.T) {
		mockRepo := &mockAttendanceRepo{
			CheckInFunc: func(eid string) (Attendance, error) {
				return Attendance{}, nil
			},
		}
		service := NewAttendanceService(mockRepo)
		_, err := service.CheckIn("1")
		assert.NoError(t, err)
	})
	t.Run("Invalid-CheckIn", func(t *testing.T) {
		mockRepo := &mockAttendanceRepo{
			CheckInFunc: func(eid string) (Attendance, error) {
				return Attendance{}, errors.New("invalid")
			},
		}
		service := NewAttendanceService(mockRepo)
		_, err := service.CheckIn("1")
		assert.Error(t, err)
		assert.Equal(t, "invalid", err.Error())
	})
}

func TestCheckOut(t *testing.T) {
	t.Run("Valid-CheckOut", func(t *testing.T) {
		mockRepo := &mockAttendanceRepo{
			CheckOutFunc: func(id int) (Attendance, error) {
				return Attendance{}, nil
			},
		}
		service := NewAttendanceService(mockRepo)
		_, err := service.CheckOut(1)
		assert.NoError(t, err)
	})
	t.Run("Invalid-CheckOut", func(t *testing.T) {
		mockRepo := &mockAttendanceRepo{
			CheckOutFunc: func(id int) (Attendance, error) {
				return Attendance{}, errors.New("invalid")
			},
		}
		service := NewAttendanceService(mockRepo)
		_, err := service.CheckOut(1)
		assert.Error(t, err)
		assert.Equal(t, "invalid", err.Error())
	})
}

func TestGetMyAttendances(t *testing.T) {
	t.Run("Valid-GetMyAttendances", func(t *testing.T) {
		mockRepo := &mockAttendanceRepo{
			GetAllMeFunc: func(eid string) ([]Attendance, error) {
				return []Attendance{}, nil
			},
		}

		service := NewAttendanceService(mockRepo)
		_, err := service.GetMyAttendances("1")
		assert.NoError(t, err)
	})
	t.Run("Invalid-GetMyAttendances", func(t *testing.T) {
		mockRepo := &mockAttendanceRepo{
			GetAllMeFunc: func(eid string) ([]Attendance, error) {
				return []Attendance{}, errors.New("invalid")
			},
		}

		service := NewAttendanceService(mockRepo)
		_, err := service.GetMyAttendances("1")
		assert.Error(t, err)
		assert.Equal(t, "invalid", err.Error())
	})
}

func TestCheckToday(t *testing.T) {
	t.Run("Valid-CheckToday", func(t *testing.T) {
		mockRepo := &mockAttendanceRepo{
			CheckTodayFunc: func(eid string) (Attendance, error) {
				return Attendance{}, nil
			},
		}

		service := NewAttendanceService(mockRepo)
		_, err := service.CheckToday("1")
		assert.NoError(t, err)
	})
	t.Run("Invalid-CheckToday", func(t *testing.T) {
		mockRepo := &mockAttendanceRepo{
			CheckTodayFunc: func(eid string) (Attendance, error) {
				return Attendance{}, errors.New("invalid")
			},
		}

		service := NewAttendanceService(mockRepo)
		_, err := service.CheckToday("1")
		assert.Error(t, err)
		assert.Equal(t, "invalid", err.Error())
	})
}
