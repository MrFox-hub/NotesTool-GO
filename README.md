# Notes Tool Documentation

The Notes Tool is a simple command-line application for managing personal notes. It allows users to add, view, and delete notes in a specified collection. This documentation provides an overview of its functionalities, usage, and examples.

## 1. Overview

The Notes Tool enables users to handle short, single-line notes. Notes can be encrypted or stored in plain text, and each note is timestamped. The tool supports basic operations such as viewing, adding, and deleting notes.

### Key Features
- **Note Management:** Add, view, and delete notes.
- **Encryption:** Optionally encrypt notes using a basic Vigenère cipher.
- **Timestamping:** Automatically timestamp each note upon creation.
- **Color Coding:** Provides a color-coded display for improved readability.

### How It Works
- **Encryption:** Uses a fixed Vigenère cipher key to encrypt and decrypt note contents.
- **File Handling:** Notes are stored in a file specified by the user. The file is loaded and saved as needed.

## 2. Example Usage

### Running the Tool
To start the tool, use the following command:

```bash
./notestool [COLLECTION_NAME]
```

Replace `[COLLECTION_NAME]` with the name of the file where notes will be stored.

#### Example

```bash
$ ./notestool mynotes
```

### Main Menu Options
Once you start the tool, you'll see the main menu with the following options:
1. **Show notes:** Displays all notes in the collection.
2. **Add a note:** Prompts the user to enter a note name and content. Optionally, the note can be encrypted.
3. **Delete a note:** Allows the user to delete a specific note by its index.
4. **Exit:** Exits the application.

### Example Workflows

#### Adding a Note
1. **Command:** Start the tool by running `./notestool mynotes`.

    ```bash
    $ ./notestool mynotes
    ```

2. **Select Operation:** Choose `2` to add a new note.

    ```plaintext
    Select operation:
    1. Show notes.
    2. Add a note.
    3. Delete a note.
    4. Exit.

    Your choice: 2
    ```

3. **Enter Note Name:** Type the note's name.

    ```plaintext
    Enter the note name:
    Meeting
    ```

4. **Enter Note Content:** Type the content of the note.

    ```plaintext
    Enter the note content:
    Discuss project updates
    ```

5. **Encrypt Note?** Choose whether to encrypt the note. Here, `n` for no.

    ```plaintext
    Do you want to encrypt this note? (y/n):
    n
    ```

6. **Result:** The note is added and saved with the current timestamp.

    ```plaintext
    Non-encrypted note added successfully.
    ```

#### Viewing Notes
1. **Command:** Start the tool by running `./notestool mynotes`.

    ```bash
    $ ./notestool mynotes
    ```

2. **Select Operation:** Choose `1` to show notes.

    ```plaintext
    Select operation:
    1. Show notes.
    2. Add a note.
    3. Delete a note.
    4. Exit.

    Your choice: 1
    ```

3. **Result:** Displays all notes with their names, timestamps, and content.

    ```plaintext
    Notes:
    001 - [Meeting] Discuss project updates [2024-08-18 14:35:22]
    ```

#### Deleting a Note
1. **Command:** Start the tool by running `./notestool mynotes`.

    ```bash
    $ ./notestool mynotes
    ```

2. **Select Operation:** Choose `3` to delete a note.

    ```plaintext
    Select operation:
    1. Show notes.
    2. Add a note.
    3. Delete a note.
    4. Exit.

    Your choice: 3
    ```

3. **Select Note Number:** Enter the number of the note you want to delete.

    ```plaintext
    Notes:
    001 - [Meeting] Discuss project updates [2024-08-18 14:35:22]

    Enter the number of the note to remove or 0 to cancel:
    1
    ```

4. **Result:** The selected note is removed from the collection.

    ```plaintext
    Note deleted successfully.
    ```

### 3. Encryption Details

#### Vigenère Cipher

The Vigenère Cipher is used for encrypting and decrypting note content. It utilizes a keyword to determine the shift for each letter in the plaintext.

#### Cipher Details
- **Key:** Fixed keyword `"KOODJOHVI"`.
- **Mapping:** Each letter in the plaintext is shifted by the position of the corresponding letter in the keyword.

#### How It Works
- **Encryption:** 
  1. Convert each letter of the plaintext to its numerical value.
  2. Add the corresponding letter's value from the keyword.
  3. Wrap around the alphabet if necessary.
  4. Convert the result back to a letter.
- **Decryption:** Reverses the encryption process using the same keyword.

#### Example
- **Plaintext:** `"Discuss"`
- **Keyword:** `"KOODJOHVI"`
- **Encrypted Message:** `"Fuhvww"`

Each letter in `"Discuss"` is shifted according to the keyword.

---

### 4. Summary

This documentation provides an overview of the Notes Tool, explaining how to use the application for managing notes, including the encryption feature. It helps users effectively handle their notes with options to view, add, and delete them.

This Tool is for localuse and one of my first projects

---