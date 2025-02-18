package postprocessing

import (
	"time"

	user "github.com/cs3org/go-cs3apis/cs3/identity/user/v1beta1"
	provider "github.com/cs3org/go-cs3apis/cs3/storage/provider/v1beta1"
	"github.com/cs3org/reva/v2/pkg/events"
)

// Postprocessing handles postprocessing of a file
type Postprocessing struct {
	ID         string
	URL        string
	User       *user.User
	Filename   string
	Filesize   uint64
	ResourceID *provider.ResourceId
	Steps      []events.Postprocessingstep
	PPDelay    time.Duration
}

// New returns a new postprocessing instance
func New(uploadID string, uploadURL string, user *user.User, filename string, filesize uint64, resourceID *provider.ResourceId, steps []events.Postprocessingstep, delay time.Duration) *Postprocessing {
	return &Postprocessing{
		ID:         uploadID,
		URL:        uploadURL,
		User:       user,
		Filename:   filename,
		Filesize:   filesize,
		ResourceID: resourceID,
		Steps:      steps,
		PPDelay:    delay,
	}
}

// Init is the first step of the postprocessing
func (pp *Postprocessing) Init(ev events.BytesReceived) interface{} {
	if len(pp.Steps) == 0 {
		return pp.finished(events.PPOutcomeContinue)
	}

	return pp.nextStep(pp.Steps[0])
}

// NextStep returns the next postprocessing step
func (pp *Postprocessing) NextStep(ev events.PostprocessingStepFinished) interface{} {
	switch ev.Outcome {
	case events.PPOutcomeContinue:
		return pp.next(ev.FinishedStep)
	default:
		return pp.finished(ev.Outcome)

	}
}

// Delay will sleep the configured time then continue
func (pp *Postprocessing) Delay(ev events.StartPostprocessingStep) interface{} {
	time.Sleep(pp.PPDelay)
	return pp.next(events.PPStepDelay)
}

func (pp *Postprocessing) next(current events.Postprocessingstep) interface{} {
	l := len(pp.Steps)
	for i, s := range pp.Steps {
		if s == current && i+1 < l {
			return pp.nextStep(pp.Steps[i+1])
		}
	}
	return pp.finished(events.PPOutcomeContinue)
}

func (pp *Postprocessing) nextStep(next events.Postprocessingstep) events.StartPostprocessingStep {
	return events.StartPostprocessingStep{
		UploadID:      pp.ID,
		URL:           pp.URL,
		ExecutingUser: pp.User,
		Filename:      pp.Filename,
		Filesize:      pp.Filesize,
		ResourceID:    pp.ResourceID,
		StepToStart:   next,
	}
}

func (pp *Postprocessing) finished(outcome events.PostprocessingOutcome) events.PostprocessingFinished {
	return events.PostprocessingFinished{
		UploadID:      pp.ID,
		ExecutingUser: pp.User,
		Filename:      pp.Filename,
		Outcome:       outcome,
	}
}
