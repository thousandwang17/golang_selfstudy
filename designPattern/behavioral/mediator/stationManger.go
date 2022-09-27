package main

type stationMannger struct {
	trianQueue     []tainer
	isPlatformFree bool
}

func NewstationMannger() meditorer {
	return &stationMannger{
		trianQueue:     []tainer{},
		isPlatformFree: true,
	}
}

func (s *stationMannger) canArrived(t tainer) bool {

	if s.isPlatformFree {
		s.isPlatformFree = false
		return true
	}

	s.trianQueue = append(s.trianQueue, t)
	return false
}

func (s *stationMannger) notifyAboutDepature() {
	if !s.isPlatformFree {
		s.isPlatformFree = true
	}

	if len(s.trianQueue) > 0 {
		s.trianQueue[0].permitArrival()
		s.trianQueue = s.trianQueue[1:]
	}
}
