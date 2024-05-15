package adapter

import (
	"github.com/night-sornram/employee-management/repository"
	"gorm.io/gorm"
)

type GormAdapter struct {
	db *gorm.DB
}

func NewGormAdapter(db *gorm.DB) repository.LeaveRepository {
	return &GormAdapter{
		db: db,
	}
}

func (g *GormAdapter) GetAll() ([]repository.Leave, error) {
	var leaves []repository.Leave
	if err := g.db.Find(&leaves).Error; err != nil {
		return nil, err
	}
	return leaves, nil
}

func (g *GormAdapter) GetByID(id int) (repository.Leave, error) {
	var leave repository.Leave
	if err := g.db.First(&leave, id).Error; err != nil {
		return leave, err
	}
	return leave, nil
}

func (g *GormAdapter) Create(leave repository.Leave) (repository.Leave, error) {
	if err := g.db.Create(&leave).Error; err != nil {
		return leave, err
	}
	return leave, nil
}

func (g *GormAdapter) Update(id int, leave repository.Leave) (repository.Leave, error) {
	if err := g.db.Model(&leave).Where("id = ?", id).Updates(leave).Error; err != nil {
		return leave, err
	}
	return leave, nil
}

func (g *GormAdapter) Delete(id int) error {
	if err := g.db.Delete(&repository.Leave{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (g *GormAdapter) UpdateStatus(id int, leaveStatus repository.LeaveStatus) (repository.Leave, error) {
	newLeave := repository.Leave{
		Status: leaveStatus.Status,
	}
	if err := g.db.Model(&newLeave).Where("id = ?", id).Updates(newLeave).Error; err != nil {
		return newLeave, err
	}
	return newLeave, nil
}
