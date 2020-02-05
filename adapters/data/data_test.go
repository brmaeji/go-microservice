package data

import (
	"microservice/models"
	"testing"
)

var memApt *MemoryAdapter

//TestNewAdapter validates the creation of a new Adapter from data package
func TestNewAdapter(t *testing.T) {
	adapterTypes := []AdapterType{
		0,
		1,
	}

	for _, at := range adapterTypes {
		adapter, err := NewAdapter(at)

		switch at {
		case Undefined:
			if err == nil {
				t.Errorf("expected to fail initialization of adaptery with type %v.", at)
			}
			t.Log("OK - got error when trying to create undefined type adapter")
			break
		case Memory:
			if err != nil {
				t.Errorf("could not create adapter of type %v, err: %v\n", at, err)
			}

			switch v := adapter.(type) {
			case *MemoryAdapter:
				t.Log("OK - created memory adapter successfully")
				memApt = adapter.(*MemoryAdapter)
			default:
				t.Errorf("expected adapter of type MemoryAdapter when using type %v, got %v\n", at, v)
			}
			break
		}
	}
}

func TestFind1(t *testing.T) {
	if memApt == nil {
		t.Error(" memory adapter is null!")
	}

	mentions, err := memApt.Find("The Beatles")
	if err != nil {
		t.Errorf("couldn't perform search: %v\n", err)
	}
	if len(mentions) > 0 {
		t.Errorf("couldn't perform search")
	}

	t.Logf("OK - 0 results expected, %v found", len(mentions))
}

func TestCreate(t *testing.T) {
	if memApt == nil {
		t.Error(" memory adapter is null!")
	}

	err := memApt.Create("The Beatles")
	if err != nil {
		t.Errorf("couldn't perform create: %v\n", err)
	}

	t.Log("OK - create had no errors")
}

func TestCreate2(t *testing.T) {
	if memApt == nil {
		t.Error(" memory adapter is null!")
	}

	err := memApt.Create("The Cure")
	if err != nil {
		t.Errorf("couldn't perform create2: %v\n", err)
	}

	t.Log("OK - create2 had no errors")
}

func TestFind2(t *testing.T) {
	if memApt == nil {
		t.Error(" memory adapter is null!")
	}

	mentions, err := memApt.Find("The Beatles")
	if err != nil {
		t.Errorf("couldn't perform search: %v\n", err)
	}
	if len(mentions) != 1 {
		t.Errorf("expected 1 results, got %v\n", len(mentions))
	}

	t.Logf("OK - 1 results expected, %v found", len(mentions))
}

func TestFind3(t *testing.T) {
	if memApt == nil {
		t.Error(" memory adapter is null!")
	}

	mentions, err := memApt.Find("the")
	if err != nil {
		t.Errorf("couldn't perform search: %v\n", err)
	}
	if len(mentions) != 2 {
		t.Errorf("expected 2 results, got %v\n", len(mentions))
	}

	t.Logf("OK - 2 results expected, %v found", len(mentions))
}

func TestIncrease(t *testing.T) {
	if memApt == nil {
		t.Error(" memory adapter is null!")
	}

	bm := &models.BandMention{
		Name:     "The Beatles",
		Mentions: 1,
	}

	changedBm, err := memApt.Increase(bm)
	if err != nil {
		t.Errorf("couldn't perform increase: %v\n", err)
	}
	mentions := changedBm.Mentions
	if mentions != 2 {
		t.Errorf("expected 2 mentions, got %v\n", mentions)
	}

	t.Logf("OK - 2 mentions expected, got %v\n", mentions)
}
