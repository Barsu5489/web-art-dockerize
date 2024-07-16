# Ascii Art Web Dockerize

Ascii Art Web dockerize is a golang web application that converts an input text to generate an ASCII  via web browser based on the three banner files provided in this project.

The project further utilized `docker image` and `engine` to run the project based on a `docker file` to create a `docker image` to run the project.

## Prerequisites

Make sure you have Docker installed on your system. If not follow the instructions provided on [dockers official website](https://docs.docker.com/engine/install) to install it according to the system that you have.

## Installation and setup

1. Clone this repository

```bash
git clone https://learn.zone01kisumu.ke/git/abrakingoo/ascii-art-web-dockerize.git
cd ascii-art-web-dockerize
```

2. Building the docker image

```bash
docker build . -t ascii-art-web-dockerize:latest
```

`.` means to build a docker container based on the `docker configuration`  in the current directory.

`-t` means `tag` or the name we've given the container supposed to be built.

3. Running the container

```bash
docker run -p 80:8080 ascii-art-web-dockerize:latest
```

`docker run`: Starts a new container.

`-p 80:8080`: Maps port 80 on your computer to port 8080 inside the container.

`ascii-art-web-dockerize:latest`: Uses the specified Docker image (with the latest version) to create the container.

4. Open your prefered browser and enter `http://localhost` to view the application

```bash
http://localhost
```

**Alternatively**

Alternatively withing the project a bash script named `builder.sh` has been created to automate the process faster.

To execute the script run :

```bash
sh builder.sh
```

The builder script also runs on a different port as opposed to the earlier one.

```textile
http://localhost:8080
```

**Cleaning up**

To cleanup the created docker containers creted run `cleaner.sh` script to automate the process.

```bash
sh cleaner.sh
```

Enter your text in the textarea, choose a banner style, and click "Createart" to generate ASCII art based on your input.

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
