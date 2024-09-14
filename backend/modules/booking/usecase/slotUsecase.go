package usecase

import (
	"context"
	"errors"
	"fmt"
	"main/modules/booking"
	"main/modules/booking/repository"
	"main/pkg/utils"
	"time"
)

type (
	SlotUsecaseService interface {
		InsertSlot(ctx context.Context, startTime, endTime string) (*booking.Slot, error)
		UpdateSlot(ctx context.Context, slotId string, startTime, endTime string) (*booking.Slot, error)
		FindOneSlot(ctx context.Context, slotId string) (*booking.Slot, error)
		FindAllSlots(ctx context.Context) ([]booking.Slot, error)
		EnableOrDisableSlot(ctx context.Context, slotId string, status int) (*booking.Slot, error)
		CheckSlotOverlap(ctx context.Context, startTime, endTime time.Time) (bool, error)
	}

	slotUsecase struct {
		slotRepository repository.SlotRepositoryService
	}
)

func NewSlotUsecase(slotRepo repository.SlotRepositoryService) SlotUsecaseService {
	return &slotUsecase{slotRepository: slotRepo}
}

func (u *slotUsecase) CheckSlotOverlap(ctx context.Context, startTime, endTime time.Time) (bool, error) {
    slots, err := u.slotRepository.FindAllSlots(ctx)
    if err != nil {
        return false, err
    }

    for _, slot := range slots {
        slotStartTime := utils.ParseTimeOnly(slot.StartTime).ToTime()

        slotEndTime := utils.ParseTimeOnly(slot.EndTime).ToTime()

        if (startTime.Before(slotStartTime)) && endTime.After(slotEndTime) ||
           (startTime.Equal(slotStartTime) && endTime.Equal(slotEndTime)) {
            return true, nil
        }
    }

    return false, nil
}



func (u *slotUsecase) InsertSlot(ctx context.Context, startTime, endTime string) (*booking.Slot, error) {
    hasOverlap, err := u.CheckSlotOverlap(ctx, utils.ParseTimeOnly(startTime).ToTime(), utils.ParseTimeOnly(endTime).ToTime())
    if err != nil {
        return nil, fmt.Errorf("failed to check slot overlap: %w", err)
    }
    if hasOverlap {
        return nil, errors.New("a slot with the same start and end time already exists")
    }

    // Construct the slot with time values
    slot := &booking.Slot{
        StartTime: startTime,  // Use string "HH:mm"
        EndTime:   endTime,    // Use string "HH:mm"
        Status:    1,  // Enabled by default
        CreatedAt: time.Now(),
        UpdatedAt: time.Now(),
    }

    return u.slotRepository.InsertSlot(ctx, slot)
}




func (u *slotUsecase) UpdateSlot(ctx context.Context, slotId string, startTime, endTime string) (*booking.Slot, error) {
	slot, err := u.slotRepository.FindOneSlot(ctx, slotId)
	if err != nil {
		return nil, fmt.Errorf("error: failed to find slot: %w", err)
	}

	slot.StartTime = slot.StartTime
	slot.EndTime = slot.StartTime
	slot.UpdatedAt = time.Now()

	return u.slotRepository.UpdateSlot(ctx, slot)
}

func (u *slotUsecase) FindOneSlot(ctx context.Context, slotId string) (*booking.Slot, error) {
	return u.slotRepository.FindOneSlot(ctx, slotId)
}

func (u *slotUsecase) FindAllSlots(ctx context.Context) ([]booking.Slot, error) {
	return u.slotRepository.FindAllSlots(ctx)
}

func (u *slotUsecase) EnableOrDisableSlot(ctx context.Context, slotId string, status int) (*booking.Slot, error) {
	return u.slotRepository.EnableOrDisableSlot(ctx, slotId, status)
}