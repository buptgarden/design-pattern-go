package mediator

type StationManager struct {
	IsPlatFormFree bool
	TrainQueue     []Train
}

func NewStationManager() *StationManager {
	return &StationManager{
		IsPlatFormFree: true,
	}
}

func (s *StationManager) CanArrive(t Train) bool {
	if s.IsPlatFormFree {
		s.IsPlatFormFree = false
		return true
	}
	s.TrainQueue = append(s.TrainQueue, t)
	return false
}

func (s *StationManager) NotifyAboutDeparture() {
	if !s.IsPlatFormFree {
		s.IsPlatFormFree = true
	}

	if len(s.TrainQueue) > 0 {
		firstTrainInQueue := s.TrainQueue[0]
		s.TrainQueue = s.TrainQueue[1:]
		firstTrainInQueue.PermitArrival()

	}

}
