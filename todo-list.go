package main

import "fmt"

const NMAX = 100

type Task struct {
	id          int
	title       string
	description string
	completed   bool
	priority    int
}

type tabtask [NMAX]Task

var tasks tabtask
var taskCount int = 0

func main() {
	var exit bool
	for !exit {
		showMenu()
		var choice int
		fmt.Print("Pilih menu: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			addTask()
		case 2:
			viewTasks()
		case 3:
			editTask()
		case 4:
			deleteTask()
		case 5:
			searchTask()
		case 6:
			sortTasks()
		case 7:
			fmt.Println("Terima kasih telah menggunakan aplikasi!")
			exit = true
		default:
			fmt.Println("Pilihan tidak valid!")
		}
	}
}

func showMenu() {
	fmt.Println("\n=== Aplikasi To-Do List ===")
	fmt.Println("1. Tambah Tugas")
	fmt.Println("2. Lihat Daftar Tugas")
	fmt.Println("3. Edit Tugas")
	fmt.Println("4. Hapus Tugas")
	fmt.Println("5. Cari Tugas")
	fmt.Println("6. Urutkan Tugas")
	fmt.Println("7. Keluar")
}

func addTask() {
	if taskCount < NMAX {
		var title, description string
		var priority int
		var validPriority bool
		
		fmt.Print("Judul tugas: ")
		fmt.Scan(&title)
		
		fmt.Print("Deskripsi: ")
		fmt.Scan(&description)
		
		for !validPriority {
			fmt.Print("Prioritas (1-5): ")
			fmt.Scan(&priority)
			if priority >= 1 && priority <= 5 {
				validPriority = true
			} else {
				fmt.Println("Prioritas harus antara 1-5!")
			}
		}
		
		tasks[taskCount] = Task{
			id:          taskCount + 1,
			title:       title,
			description: description,
			completed:   false,
			priority:    priority,
		}
		taskCount++
		fmt.Println("Tugas berhasil ditambahkan!")
	} else {
		fmt.Println("Daftar tugas penuh!")
	}
}

func viewTasks() {
	if taskCount > 0 {
		fmt.Println("\nDaftar Tugas:")
		for i := 0; i < taskCount; i++ {
			status := "Belum selesai"
			if tasks[i].completed {
				status = "Selesai"
			}
			fmt.Printf("%d. [%s] %s (Prioritas: %d)\n   Deskripsi: %s\n", 
				tasks[i].id, status, tasks[i].title, tasks[i].priority, tasks[i].description)
		}
	} else {
		fmt.Println("Tidak ada tugas!")
	}
}

func editTask() {
	var id int
	fmt.Print("Masukkan ID tugas yang akan diedit: ")
	fmt.Scan(&id)

	index := binarySearch(id)
	if index != -1 {
		var title, description, completedInput string
		var priority int
		var completed bool

		fmt.Printf("Judul (%s): ", tasks[index].title)
		fmt.Scan(&title)
		
		fmt.Printf("Deskripsi (%s): ", tasks[index].description)
		fmt.Scan(&description)
		
		fmt.Printf("Selesai (y/n) (%t): ", tasks[index].completed)
		fmt.Scan(&completedInput)
		
		fmt.Printf("Prioritas (%d): ", tasks[index].priority)
		fmt.Scan(&priority)

		if title != "" {
			tasks[index].title = title
		}
		if description != "" {
			tasks[index].description = description
		}
		if completedInput == "y" {
			completed = true
		} else if completedInput == "n" {
			completed = false
		} else {
			completed = tasks[index].completed
		}
		tasks[index].completed = completed
		
		if priority >= 1 && priority <= 5 {
			tasks[index].priority = priority
		}

		fmt.Println("Tugas berhasil diupdate!")
	} else {
		fmt.Println("Tugas tidak ditemukan!")
	}
}

func deleteTask() {
	var id int
	fmt.Print("Masukkan ID tugas yang akan dihapus: ")
	fmt.Scan(&id)

	index := binarySearch(id)
	if index != -1 {
		for i := index; i < taskCount-1; i++ {
			tasks[i] = tasks[i+1]
		}
		taskCount--
		fmt.Println("Tugas berhasil dihapus!")
	} else {
		fmt.Println("Tugas tidak ditemukan!")
	}
}

func searchTask() {
	var keyword string
	fmt.Print("Masukkan kata kunci pencarian: ")
	fmt.Scan(&keyword)

	found := false
	for i := 0; i < taskCount; i++ {
		if contains(tasks[i].title, keyword) || 
		   contains(tasks[i].description, keyword) {
			status := "Belum selesai"
			if tasks[i].completed {
				status = "Selesai"
			}
			fmt.Printf("%d. [%s] %s (Prioritas: %d)\n", 
				tasks[i].id , status, tasks[i].title, tasks[i].priority)
			found = true
		}
	}

	if !found {
		fmt.Println("Tidak ditemukan tugas dengan kata kunci tersebut")
	}
}

func contains(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

func binarySearch(id int) int {
	insertionSortById()

	left, right := 0, taskCount-1
	for left <= right {
		mid := (left + right) / 2
		if tasks[mid].id == id {
			return mid
		} else if tasks[mid].id < id {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func insertionSortById() {
	for i := 1; i < taskCount; i++ {
		key := tasks[i]
		j := i - 1

		for j >= 0 && tasks[j].id > key.id {
			tasks[j+1] = tasks[j]
			j = j - 1
		}
		tasks[j+1] = key
	}
}

func sortTasks() {
	fmt.Println("\n1. Urutkan menaik berdasarkan ID")
	fmt.Println("2. Urutkan menurun berdasarkan ID")
	fmt.Println("3. Urutkan menaik berdasarkan Prioritas")
	fmt.Println("4. Urutkan menurun berdasarkan Prioritas")
	fmt.Print("Pilih: ")

	var choice int
	fmt.Scan(&choice)

	switch choice {
	case 1:
		insertionSortById()
		viewTasks()
	case 2:
		selectionSortByIdDesc()
		viewTasks()
	case 3:
		selectionSortByPriorityAsc()
		viewTasks()
	case 4:
		selectionSortByPriorityDesc()
		viewTasks()
	default:
		fmt.Println("Pilihan tidak valid!")
	}
}

func selectionSortByIdDesc() {
	for i := 0; i < taskCount-1; i++ {
		maxIdx := i
		for j := i + 1; j < taskCount; j++ {
			if tasks[j].id > tasks[maxIdx].id {
				maxIdx = j
			}
		}
		tasks[i], tasks[maxIdx] = tasks[maxIdx], tasks[i]
	}
}

func selectionSortByPriorityAsc() {
	for i := 0; i < taskCount-1; i++ {
		minIdx := i
		for j := i + 1; j < taskCount; j++ {
			if tasks[j].priority < tasks[minIdx].priority {
				minIdx = j
			}
		}
		tasks[i], tasks[minIdx] = tasks[minIdx], tasks[i]
	}
}

func selectionSortByPriorityDesc() {
	for i := 0; i < taskCount-1; i++ {
		maxIdx := i
		for j := i + 1; j < taskCount; j++ {
			if tasks[j].priority > tasks[maxIdx].priority {
				maxIdx = j
			}
		}
		tasks[i], tasks[maxIdx] = tasks[maxIdx], tasks[i]
	}
}