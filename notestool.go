package main

import (
	"bufio"   // Package for buffered I/O operations.
	"fmt"     // Package for formatted I/O operations.
	"os"      // Package for interacting with the operating system.
	"os/exec" // Package for running external commands.
	"strings" // Package for manipulating string values.
	"time"    // Package for handling time-related operations.
)

// Note represents a single note with its metadata.
type Note struct {
	Name        string // Name of the note.
	IsEncrypted bool   // Indicates whether the note is encrypted.
	Timestamp   string // Timestamp when the note was created.
	Content     string // Content of the note.
}

const (
	colorReset  = "\033[0m"  // ANSI escape code to reset color.
	colorGreen  = "\033[32m" // ANSI escape code for green color.
	colorYellow = "\033[33m" // ANSI escape code for yellow color.
	colorCyan   = "\033[36m" // ANSI escape code for cyan color.
)

func main() {
	// Check if the program was called with exactly one argument or if "help" is requested.
	if len(os.Args) != 2 || os.Args[1] == "help" {
		printHelp() // Print help information.
		return      // Exit the program.
	}

	collectionName := os.Args[1]       // Get the collection name from the command-line argument.
	notes := loadNotes(collectionName) // Load existing notes from the specified file.

	clearScreen()                             // Clear the console screen.
	fmt.Println("Welcome to the notes tool!") // Welcome message.

	werePurged := true // Flag to check if any notes were purged.
	if len(notes) == 0 {
		werePurged = false // No notes found, so nothing was purged.
	}

	// Main loop for handling user interactions.
	for {
		fmt.Println("\nSelect operation:")
		fmt.Println("1. Show notes.")
		fmt.Println("2. Add a note.")
		fmt.Println("3. Delete a note.")
		fmt.Println("4. Exit.")

		fmt.Print("\nYour choice: ")
		var choice int
		_, err := fmt.Scanf("%d\n", &choice) // Get user's choice as an integer.
		if err != nil {
			clearScreen()
			fmt.Println("Invalid input, please enter a number between 1 and 4.")
			continue // Restart the loop on invalid input.
		}

		// Handle user's choice.
		switch choice {
		case 1:
			clearScreen()
			showNotes(notes) // Display all notes.
		case 2:
			clearScreen()
			addNote(&notes, &werePurged)     // Add a new note.
			saveNotes(collectionName, notes) // Save updated notes to the file.
		case 3:
			clearScreen()
			deleteNote(&notes, &werePurged)  // Delete a note.
			saveNotes(collectionName, notes) // Save updated notes to the file.
		case 4:
			clearScreen()
			fmt.Println("Thank you for using notes tool!\n\nExiting...")
			return // Exit the program.
		default:
			clearScreen()
			fmt.Println("Invalid choice, please enter a number between 1 and 4.")
		}
	}
}

// printHelp displays usage instructions for the program.
func printHelp() {
	fmt.Println("\nUsage: ./notestool [COLLECTION_NAME]")
	fmt.Println("Manage short single-line notes in the specified collection.")
}

// loadNotes reads notes from a file and returns them as a slice of Note structs.
func loadNotes(filename string) []Note {
	file, err := os.Open(filename) // Open the specified file for reading.
	if err != nil {
		if os.IsNotExist(err) {
			return []Note{} // Return an empty slice if the file does not exist.
		}
		clearScreen()
		fmt.Println("Error opening file:", err)
		os.Exit(1) // Exit the program on any other error.
	}
	defer file.Close() // Ensure the file is closed when the function exits.

	var notes []Note
	scanner := bufio.NewScanner(file) // Create a new scanner for the file.
	for scanner.Scan() {
		line := scanner.Text()                // Read a line from the file.
		parts := strings.SplitN(line, ":", 4) // Split the line into 4 parts.
		if len(parts) == 4 {
			isEncrypted := parts[1] == "true" // Determine if the note is encrypted.
			notes = append(notes, Note{Name: parts[0], IsEncrypted: isEncrypted, Timestamp: parts[2], Content: parts[3]})
		}
	}

	if err := scanner.Err(); err != nil {
		clearScreen()
		fmt.Println("Error reading file:", err)
		os.Exit(1) // Exit the program if there was an error reading the file.
	}

	return notes // Return the loaded notes.
}

// saveNotes writes the current notes to the specified file.
func saveNotes(filename string, notes []Note) {
	file, err := os.Create(filename) // Create a new file for writing.
	if err != nil {
		clearScreen()
		fmt.Println("Error creating file:", err)
		os.Exit(1) // Exit the program if the file cannot be created.
	}
	defer file.Close() // Ensure the file is closed when the function exits.

	writer := bufio.NewWriter(file) // Create a buffered writer.
	for _, note := range notes {
		_, err := writer.WriteString(note.Name + ":" + fmt.Sprint(note.IsEncrypted) + ":" + note.Timestamp + ":" + note.Content + "\n")
		if err != nil {
			clearScreen()
			fmt.Println("Error writing to file:", err)
			os.Exit(1) // Exit the program if there is an error writing to the file.
		}
	}
	writer.Flush() // Flush the writer to ensure all data is written.
}

// showNotes displays all the notes to the user.
func showNotes(notes []Note) {
	if len(notes) == 0 {
		clearScreen()
		fmt.Println("No notes available.") // Inform the user if no notes are available.
		return
	}

	fmt.Println("Notes:")
	for i, note := range notes {
		content := note.Content
		if note.IsEncrypted {
			content = decrypt(content) // Decrypt the note content if it is encrypted.
		}
		// Display the note index, name, content, and timestamp with color coding.
		fmt.Printf("%s%03d%s - %s[%s]%s %s %s[%s]%s\n", colorYellow, i+1, colorReset, colorGreen, note.Name, colorReset, content, colorCyan, note.Timestamp, colorReset)
	}
	fmt.Println()
}

// addNote prompts the user to add a new note to the collection.
func addNote(notes *[]Note, werePurged *bool) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter the note name:")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name) // Trim whitespace from the input.
	if name == "" {
		clearScreen()
		fmt.Println("Note name cannot be empty.") // Ensure the note has a name.
		return
	}

	fmt.Println("\nEnter the note content:")
	content, _ := reader.ReadString('\n')
	content = strings.TrimSpace(content) // Trim whitespace from the input.
	if content == "" {
		clearScreen()
		fmt.Println("Note content cannot be empty.") // Ensure the note has content.
		return
	}

	fmt.Println("\nDo you want to encrypt this note? (y/n):")
	encryptChoice, _ := reader.ReadString('\n')
	encryptChoice = strings.TrimSpace(strings.ToLower(encryptChoice)) // Get the user's choice and normalize it.

	isEncrypted := false
	if encryptChoice == "y" {
		content = encrypt(content) // Encrypt the note content if requested.
		isEncrypted = true
		clearScreen()
		fmt.Println("Encrypted note added successfully.")
	} else if encryptChoice != "n" {
		clearScreen()
		fmt.Println("Warning: invalid input was defaulted to 'n', non-encrypted note was added.")
	} else {
		clearScreen()
		fmt.Println("Non-encrypted note added successfully.")
	}

	timestamp := time.Now().Format("2006-01-02 15:04:05")                                                       // Get the current timestamp.
	*notes = append(*notes, Note{Name: name, IsEncrypted: isEncrypted, Timestamp: timestamp, Content: content}) // Add the new note.
	*werePurged = true                                                                                          // Mark that notes have been added or modified.
}

// deleteNote prompts the user to delete a note from the collection.
func deleteNote(notes *[]Note, werePurged *bool) {
	if len(*notes) == 0 {
		if *werePurged {
			clearScreen()
			fmt.Println("You already purged all of your most wildest ideas.") // Inform the user if all notes were deleted.
			return
		} else {
			clearScreen()
			fmt.Println("The collection is already empty.") // Inform the user if the collection is empty.
			return
		}
	}

	showNotes(*notes) // Display all notes for selection.

	fmt.Println("Enter the number of the note to remove or 0 to cancel:")
	var choice int
	_, err := fmt.Scanf("%d\n", &choice)
	if err != nil {
		clearScreen()
		fmt.Println("Invalid input, please enter a note index.") // Handle invalid input.
		return
	} else if choice < 0 || choice > len(*notes) {
		clearScreen()
		fmt.Println("Invalid choice, note index out of bounds.") // Ensure the choice is within valid bounds.
		return
	}

	if choice == 0 {
		clearScreen()
		fmt.Println("No notes were deleted.") // Handle case where deletion is canceled.
		return
	}

	*notes = append((*notes)[:choice-1], (*notes)[choice:]...) // Remove the selected note.
	clearScreen()
	fmt.Println("Note deleted successfully.")
}

// encrypt applies a simple encryption algorithm to the note content.
func encrypt(message string) string {
	encrypted := ""
	key := "KOODJOHVI" // Encryption key.
	keyLen := len(key)
	for i, char := range message {
		if char >= 'A' && char <= 'Z' {
			shift := int(key[i%keyLen]) - 'A'
			encrypted += string((int(char)+shift-'A')%26 + 'A') // Encrypt uppercase letters.
		} else if char >= 'a' && char <= 'z' {
			shift := int(key[i%keyLen]) - 'A'
			encrypted += string((int(char)+shift-'a')%26 + 'a') // Encrypt lowercase letters.
		} else {
			encrypted += string(char) // Leave non-alphabetic characters unchanged.
		}
	}
	return encrypted
}

// decrypt reverses the simple encryption algorithm to get the original content.
func decrypt(message string) string {
	decrypted := ""
	key := "KOODJOHVI" // Same key used for decryption.
	keyLen := len(key)
	for i, char := range message {
		if char >= 'A' && char <= 'Z' {
			shift := int(key[i%keyLen]) - 'A'
			decrypted += string((int(char)-shift+26-'A')%26 + 'A') // Decrypt uppercase letters.
		} else if char >= 'a' && char <= 'z' {
			shift := int(key[i%keyLen]) - 'A'
			decrypted += string((int(char)-shift+26-'a')%26 + 'a') // Decrypt lowercase letters.
		} else {
			decrypted += string(char) // Leave non-alphabetic characters unchanged.
		}
	}
	return decrypted
}

// clearScreen clears the console screen using the "clear" command.
func clearScreen() {
	cmd := exec.Command("clear") // Create a command to clear the screen.
	cmd.Stdout = os.Stdout       // Set the command's output to the console.
	cmd.Run()                    // Execute the command.
}
