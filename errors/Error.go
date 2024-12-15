package errors

import "errors"

var (
	ErrorOpenFile          = errors.New("Не получилось открыт файл!\n")
	ErrorDecoderFile       = errors.New("Не получилоьс прочитать файл!")
	ErrrorRunBot           = errors.New("Не получлось запустить бота!\n")
	ErrorLoadConfiguration = errors.New("Ошибка загрузки конфигурации:")
)
