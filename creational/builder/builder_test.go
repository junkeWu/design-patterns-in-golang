package builder

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMilkTeaBuilder_Build(t *testing.T) {
	buildMilkTea("巧克力奶茶","大杯",false ,false)
	buildMilkTea("巧克力奶昔","中杯",true ,false)
	tea := buildMilkTea("草莓奶茶", "大杯", true, true)
	assert.Equal(t, tea.Type, "草莓奶茶","Error")
	assert.Equal(t, tea.Size, "大杯","Error")
	assert.Equal(t, tea.Pearl, true,"Error")
	assert.Equal(t, tea.Ice, true,"Error")
	assert.Equal(t, tea, &MilkTea{
		"草莓奶茶",
		"大杯",
		true,
		true,
	})
}


