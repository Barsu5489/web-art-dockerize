# Ascii Art Web

Ascii Art Web is a web application that generates ASCII art representations based on user input and chosen banner styles.

## Description

Ascii Art Web is a simple web application built with Go (Golang) that allows users to input text and select from predefined banner styles to generate ASCII art. It processes the user input, handles special characters, and displays the ASCII art output using HTML and Tailwind CSS.



## Usage: How to Run

To run the project locally, follow these steps:

1. Clone the repository:

```bash
   git clone https://learn.zone01kisumu.ke/git/abrakingoo/ascii-art-web.git
   cd ascii-art-web
```

2. Start the server on port 8080:

```
go run .
```
3. Open your web browser and navigate to:

```
http://localhost:8080/

```

4. Enter your text in the textarea, choose a banner style, and click "Create art" to generate ASCII art based on your input.

# Implementation Details: Algorithm
- The core functionality of Ascii-art-Web involves converting user-provided text into ASCII art. Here's an overview of the implementation:

### 1. Server Setup:

- The application runs a web server using Go's net/http package.

#### 2. Handling Requests:

- GET /: Serves the index.html file located in the templates directory.
- POST /: Processes form submissions containing user input for generating ASCII art.

#### 3. ASCII Art Generation:

- Reads ASCII character representations from files (standard.txt, thinkertoy.txt, shadow.txt).
- Processes user input to handle special characters like newline (\n) and tab (\t).
- Generates ASCII art by mapping characters to their respective ASCII representations.

#### 4. Error Handling:

- Checks for empty ASCII art files (standard.txt, thinkertoy.txt, shadow.txt).
- Handles issues related to newline characters in the input string.
<br>
<br>

This application is a basic implementation of ASCII art generation using Go and Tailwind CSS, <br>suitable for educational purposes or as a starting point for more complex ASCII art applications.

### Authors
<table>
<tr>
    <td align="center" style="word-wrap: break-word; width: 150.0; height: 150.0">
        <a href=https://learn.zone01kisumu.ke/git/aaochieng>
            <img src=https://learn.zone01kisumu.ke/git/avatars/8a1b24358854eb12998a07c269542193?size=870 width="100;"  style="border-radius:50%;align-items:center;justify-content:center;overflow:hidden;padding-top:10px" alt=Aaron/>
            <br />
            <sub style="font-size:14px"><b>Aoron Ochieng</b></sub>
        </a>
    </td>
    <td align="center" style="word-wrap: break-word; width: 150.0; height: 150.0">
        <a href=https://learn.zone01kisumu.ke/git/ebarsula>
            <img src=https://learn.zone01kisumu.ke/git/avatars/fa966ef34b0ccdfe772414745aeee49f?size=870 width="100;"  style="border-radius:50%;align-items:center;justify-content:center;overflow:hidden;padding-top:10px" alt=Emmanuel/>
            <br />
            <sub style="font-size:14px"><b>Emmanuel Barsulai</b></sub>
        </a>
    </td>
    <td align="center" style="word-wrap: break-word; width: 150.0; height: 150.0">
        <a href=https://learn.zone01kisumu.ke/git/abrakingoo>
            <img src=https://learn.zone01kisumu.ke/git/avatars/c307852c0cb9222c1ea2c71f98ff2d51?size=870 width="100;"  style="border-radius:50%;align-items:center;justify-content:center;overflow:hidden;padding-top:10px" alt=Abraham/>
            <br />
            <sub style="font-size:14px"><b>Abraham kingoo</b></sub>
        </a>
    </td>
</tr>
</table>

## License 
© 2024 Ascii Art Web™. All Rights Reserved.

This README.md file provides a clear overview of our project, including its purpose, how to run it, <br> details about the ASCII art generation algorithm, and credits to the authors.<br> Adjust the URLs and details as per your specific project setup and preferences.

