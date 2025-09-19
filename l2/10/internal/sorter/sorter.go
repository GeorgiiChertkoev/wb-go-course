package sorter

import (
	"runtime"
	"unix-sort/internal/args"
	"unix-sort/internal/comparator"
	"unix-sort/internal/reader"
)

// основная логика сортировки
func Sort(opts args.SortOptions) error {
	/*

		будем общаться через chan []string где каждый слайс чуть больше 1гб
		("чуть больше" потому что будем считывать пока не найдем перенос строки и размер слайса выйдет за 1 гб)
		// далее 1гб можно поменять на buffer-size из параметров

		сортим слайс и кладем его в временный файл

		больше одной горутины на файл = бессмысленно из-за syscall-ов
		одна горутина для чтения файла
		GOMAXPROCS или -parallel  горутин для сортировки и записи

		когда мы всё прочитали, канал закрывается, мы досортируем все файлы и сохраним их в времнную директорию

		дальше сделаем k-way merge. будем считывать некоторое число строк из каждого файла в свою очередь
		дальше добавим в мин кучу первый элемент каждой очереди
		и будем доставать из нее значения, заменяя их следующими из соотв очереди

	*/
	comparer, err := comparator.BuildComparator(opts)
	if err != nil {
		return err
	}

	slices := make(chan []string, runtime.GOMAXPROCS(0))

	go reader.Read(opts, slices)

	return nil

}
