package fo

import (
	"os"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestA(t *testing.T) {
	may := NewMay[time.Time]()

	_ = may.Invoke(time.Parse("2006-01-02", "bad-value"))
	_ = may.Invoke(time.Parse("2006-01-02", "bad-value2"))

	if err := may.CollectAsError(); err != nil {
		return err
	}
}
