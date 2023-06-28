Certainly! Here's an example of an MD (Markdown) file documenting the commands you provided:

```markdown
## Setup Instructions

To set up and run the web application, follow the steps below:

1. Update the package lists and upgrade installed packages:

```bash
apt update -y
```

2. Install the required packages `wget` and `zip`:

```bash
apt install -y wget zip
```

3. Download the webapp.zip file from the specified URL:

```bash
wget https://stackgenie-web.s3.amazonaws.com/webapp.zip
```

4. Unzip the webapp.zip file:

```bash
unzip webapp.zip
```

5. Move into the extracted directory:

```bash
cd webapp
```

6. Execute the webapp:

```bash
./webapp
```

Once you have completed these steps, the web application should be up and running.

Note: The commands assume you are using a Debian-based Linux distribution with the `apt` package manager.
```

You can save this content in a file with a `.md` extension, for example, `setup.md`. The Markdown file provides structured and formatted instructions for setting up and running the web application using the provided commands.


or you can use the following command in single step:

```bash
sudo apt update -y  && sudo apt install -y wget zip && sudo mkdir webapp && cd webapp && wget https://stackgenie-web.s3.amazonaws.com/webapp.zip && unzip webapp && ./webapp &
```