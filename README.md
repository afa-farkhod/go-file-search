# 🔍 Go File Search Utility
A simple CL tool to recursively search for a keyword in all files within a given directory


## 📦 Features

- 🔎 Search for a keyword in all files (not just `.txt`)
- 📂 Recursively walks through all subdirectories
- 📄 Prints matching lines with file names and line numbers
- ⚡ Fast and lightweight — perfect for local code or note searches

## 🚀 Usage

```
go run main.go <directory-or-file-path> <keyword>
```

## 🧪 Examples

```
# search all files in a directory:
go run main.go ./my-folder cosmos

# search a single file:
go run main.go ./go.mod cosmos
```

## 🛠 Built With

- Go `filepath.Walk`
- Go `bufio.Scanner`
- Simple string matching with `strings.Contains`

## 🧱 Project Structure
```
├── main.go         # Entry point with CLI logic
└── README.md       # Project documentation
```

## 💡 Future Improvements

- Add support for regex pattern search
- Colorized terminal output
- Save results to file

## 📜 License

[MIT License](https://github.com/afa-farkhod/go-file-search/tree/main?tab=MIT-1-ov-file)