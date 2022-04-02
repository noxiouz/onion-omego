package oled

const (
	i2COledAddr = 0x3C

	OLED_EXP_WIDTH       uint8 = 128
	OLED_EXP_HEIGHT      uint8 = 64
	OLED_EXP_PAGES       uint8 = 8
	OLED_EXP_CHAR_LENGTH uint8 = 6
	OLED_EXP_NUM_CHARS   uint8 = 96

	OLED_EXP_CHAR_COLUMNS       uint8 = 21
	OLED_EXP_CHAR_ROWS          uint8 = 8
	OLED_EXP_CHAR_COLUMN_PIXELS uint8 = (OLED_EXP_CHAR_COLUMNS * OLED_EXP_CHAR_LENGTH)

	OLED_EXP_CONTRAST_MIN uint8 = 0
	OLED_EXP_CONTRAST_MAX uint8 = 255

	OLED_EXP_DEF_CONTRAST_EXTERNAL_VCC   uint8 = 0x9f
	OLED_EXP_DEF_CONTRAST_SWITCH_CAP_VCC uint8 = 0xcf

	// Registers
	OLED_EXP_REG_DATA    uint8 = 0x40
	OLED_EXP_REG_COMMAND uint8 = 0x80

	// Addresses
	OLED_EXP_ADDR_BASE_PAGE_START uint8 = 0xB0

	// Command Constants
	OLED_EXP_SET_CONTRAST          uint8 = 0x81
	OLED_EXP_DISPLAY_ALL_ON_RESUME uint8 = 0xA4
	OLED_EXP_DISPLAY_ALL_ON        uint8 = 0xA5
	OLED_EXP_NORMAL_DISPLAY        uint8 = 0xA6
	OLED_EXP_INVERT_DISPLAY        uint8 = 0xA7
	OLED_EXP_DISPLAY_OFF           uint8 = 0xAE
	OLED_EXP_DISPLAY_ON            uint8 = 0xAF
	OLED_EXP_SET_DISPLAY_OFFSET    uint8 = 0xD3
	OLED_EXP_SET_COM_PINS          uint8 = 0xDA
	OLED_EXP_SET_VCOM_DETECT       uint8 = 0xDB
	OLED_EXP_SET_DISPLAY_CLOCK_DIV uint8 = 0xD5
	OLED_EXP_SET_PRECHARGE         uint8 = 0xD9
	OLED_EXP_SET_MULTIPLEX         uint8 = 0xA8
	OLED_EXP_SET_LOW_COLUMN        uint8 = 0x00
	OLED_EXP_SET_HIGH_COLUMN       uint8 = 0x10
	OLED_EXP_SET_START_LINE        uint8 = 0x40
	OLED_EXP_MEMORY_MODE           uint8 = 0x20

	OLED_EXP_MEM_HORIZONTAL_ADDR_MODE uint8 = 0x00
	OLED_EXP_MEM_VERTICAL_ADDR_MODE   uint8 = 0x01
	OLED_EXP_MEM_PAGE_ADDR_MODE       uint8 = 0x02
	OLED_EXP_MEM_NUM_MODES            uint8 = 3

	OLED_EXP_COLUMN_ADDR    uint8 = 0x21
	OLED_EXP_PAGE_ADDR      uint8 = 0x22
	OLED_EXP_COM_SCAN_INC   uint8 = 0xC0
	OLED_EXP_COM_SCAN_DEC   uint8 = 0xC8
	OLED_EXP_SEG_REMAP      uint8 = 0xA0
	OLED_EXP_CHARGE_PUMP    uint8 = 0x8D
	OLED_EXP_EXTERNAL_VCC   uint8 = 0x01
	OLED_EXP_SWITCH_CAP_VCC uint8 = 0x02

	// Scrolling Constants
	OLED_EXP_ACTIVATE_SCROLL                      uint8 = 0x2F
	OLED_EXP_DEACTIVATE_SCROLL                    uint8 = 0x2E
	OLED_EXP_SET_VERTICAL_SCROLL_AREA             uint8 = 0xA3
	OLED_EXP_RIGHT_HORIZONTAL_SCROLL              uint8 = 0x26
	OLED_EXP_LEFT_HORIZONTAL_SCROLL               uint8 = 0x27
	OLED_EXP_VERTICAL_AND_RIGHT_HORIZONTAL_SCROLL uint8 = 0x29
	OLED_EXP_VERTICAL_AND_LEFT_HORIZONTAL_SCROLL  uint8 = 0x2A

	OLED_EXP_SCROLL_SPEED_5_FRAMES   uint8 = 0x00
	OLED_EXP_SCROLL_SPEED_64_FRAMES  uint8 = 0x01
	OLED_EXP_SCROLL_SPEED_128_FRAMES uint8 = 0x02
	OLED_EXP_SCROLL_SPEED_256_FRAMES uint8 = 0x03
	OLED_EXP_SCROLL_SPEED_3_FRAMES   uint8 = 0x04
	OLED_EXP_SCROLL_SPEED_4_FRAMES   uint8 = 0x05
	OLED_EXP_SCROLL_SPEED_25_FRAMES  uint8 = 0x06
	OLED_EXP_SCROLL_SPEED_2_FRAMES   uint8 = 0x07
	OLED_EXP_SCROLL_SPEED_NUM        uint8 = 8

	OLED_EXP_READ_LCD_STRING_OPT0     = "0x%02x,"
	OLED_EXP_READ_LCD_STRING_OPT1     = "%2x"
	OLED_EXP_READ_LCD_DATA_IDENTIFIER = "data:"
)