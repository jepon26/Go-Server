# Go-Server
A Web Server With Go


This is a simple Go web server that handles form submissions and serves static files. It demonstrates the basic structure and functionality of a web server using Go's built-in net/http package.

## Features


Accepts form submissions with user information (name and address)
Parses form data and renders it using an HTML template
Serves static files (HTML, CSS, JavaScript, etc.) from the "static" directory


## Prerequisites

Before running the server, ensure that you have the following:

* Go programming language installed on your machine **(Installation guide)**
* A web browser to interact with the server


## Getting Started


1. Clone this repository to your local machine or download the source code as a ZIP file.

2. Navigate to the project directory in your terminal.

3. Open the **template.html** file and customize it to your liking. This file represents the template for rendering the user information.

4. (Optional) Add your static files (HTML, CSS, JavaScript, etc.) to the "static" directory.

5. Run the following command to start the server:

````shell
Copy code
go run main.go
The server should now be running on http://localhost:3000.
     `````shell


## Usage

### Submitting the Form

* Open your web browser and navigate to ***http://localhost:3000**
* Fill in the form fields with your name and address.
*Click the "Submit" button to submit the form.
* The server will render the submitted user information using the template.html file.

## Serving Static Files

* To access static files (HTML, CSS, JavaScript, etc.), simply navigate to the respective URLs on the server.
* For example, if you have an **index.html** file in the "static" directory, you can access it at **http://localhost:3000/index.html.**
Contributing




