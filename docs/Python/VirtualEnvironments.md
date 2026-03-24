# About Virtual Environments

> - Used to contain a specific Python interpreter and software libraries and binaries which are needed to support a project (library or application). These are by default isolated from software in other virtual environments and Python interpreters and libraries installed in the operating system.
> - Contained in a directory, conventionally named .venv or venv in the project directory, or under a container directory for lots of virtual environments, such as ~/.virtualenvs.
> - Not checked into source control systems such as Git.
> - Considered as disposable – it should be simple to delete and recreate it from scratch. You don’t place any project code in the environment.
> - Not considered as movable or copyable – you just recreate the same environment in the target location.

## Create venv

```bash
python3 -m venv /path/to/new/virtual/environment
```


## Activate venv

```bash
source /path/to/new/virtual/environment/bin/activate
```

## Deactivate venv

```bash
deactivate
```

## Install packages

```bash

# Install a package
pip install package_name

# Install a specific version of a package
pip install package_name==1.0.0

# Install from a requirements file

pip install -r requirements.txt
```

## Usefulness

> - *Avoid System Pollution*:
> Linux and macOS come preinstalled with a version of Python that the operating system uses for internal tasks.
>If you install packages to your operating system’s global Python, these packages will mix with the system-relevant packages. This mix-up could have unexpected side effects on tasks crucial to your operating system’s normal behavior.
> - *Sidestep Dependency Conflicts*
> ne of your projects might require a different version of an external library compared to another project. If you only have one place to install packages, then you won’t be able to work with two different versions of the same library. This is a common reason why it’s recommended to use a Python virtual environment.
> - Minimize Reproducibility Issues
> If you’re working on a project with a team, you’ll want to ensure that everyone is using the same versions of the same libraries. If you use a virtual environment, you can share the requirements file with your team, and they can install the same versions of the same libraries.
> - *Dodge Installation Privilege Lockouts*
> you may need administrator privileges on a computer to install packages into the host Python’s site-packages directory. In a corporate work environment, you most likely won’t have that level of access to the machine that you’re working on.

## What is Python Virtual Environment?

> The short answer is that a Python virtual environment is a folder structure that gives you everything you need to run a lightweight yet isolated Python environment.

### A Folder Structure

> When you create a new virtual environment using the venv module, Python creates a self-contained folder structure and copies or symlinks the Python executable files into that folder structure.

```text
venv/
│
├── bin/
│   ├── Activate.ps1
│   ├── activate
│   ├── activate.csh
│   ├── activate.fish
│   ├── pip
│   ├── pip3
│   ├── pip3.12
│   ├── python
│   ├── python3
│   └── python3.12
│
├── include/
│   │
│   └── python3.12/
│
├── lib/
│   │
│   └── python3.12/
│       │
│       └── site-packages/
│           │
│           ├── pip/
│           │
│           └── pip-24.2.dist-info/
│
├── lib64/
│   │
│   └── python3.12/
│       │
│       └── site-packages/
│           │
│           ├── pip/
│           │
│           └── pip-24.2.dist-info/
│
└── pyvenv.cfg
```

> - bin/ contains the executable files of your virtual environment. Most notable are the Python interpreter (python) and the pip executable (pip), as well as their respective symlinks (python3, python3.12, pip3, pip3.12). The folder also contains activation scripts for your virtual environment. Your specific activation script depends on what shell you use. For example, in this tutorial, you ran activate, which works for the Bash and Zsh shells.

> - include/ is an initially empty folder that Python uses to include C header files for packages you might install that depend on C extensions.

> - lib/ contains the site-packages/ directory nested in a folder that designates the Python version (python3.12/). site-packages/ is one of the main reasons for creating your virtual environment. This folder is where you’ll install the external packages that you want to use within your virtual environment. Starting with Python 3.12, your virtual environment comes preinstalled with only one dependency, pip. You’ll learn more about it in a bit.

> - lib64/ in many Linux systems comes as a symlink to lib/ for compatibility reasons. Some Linux systems may use the distinction between lib/ and lib64/ to install different versions of libraries depending on their architecture.

> - A {name}-{version}.dist-info/ directory, which you get by default for pip, contains package distribution information that exists to record information about installed packages.

> - pyvenv.cfg is a crucial file for your virtual environment. It contains only a couple of key-value pairs that Python uses to set variables in the sys module that determine which Python interpreter and site-packages directory the current Python session will use. You’ll learn more about the settings in this file when you read about how a virtual environment works.

## Some notes

> You can access Python’s standard-library modules because your virtual environment reuses Python’s built-ins and standard-library modules from the Python installation you used to create your virtual environment.

```shell
python3 -m venv venv/ --system-site-packages
```

> If you add --system-site-packages when you call venv, Python will set the value to include-system-site-packages in pyvenv.cfg to true. This setting means that you can use any external packages that you installed to your base Python as if you’d installed them into your virtual environment.

> if you are encountring this error :

```shell

error: externally-managed-environment

× This environment is externally managed
╰─> To install Python packages system-wide, try apt install
    python3-xyz, where xyz is the package you are trying to
    install.

```

> use the following command:

```shell

python3 -m pip config set global.break-system-packages true
```


# pipenv

> a simple python virtual environment

```shell

sudo apt install pipenv
pip install --user pipenv
pipenv install package_name
```

> run the code using :

```shell

pipenv run python script.py
```

## Create a virtual environment

```shell

virtualenv venv
```

> virtualenv venv will create a folder in the current directory which will contain the Python executable files, and a copy of the pip library which you can use to install other packages. The name of the virtual environment (in this case, it was venv) can be anything; omitting the name will place the files in the current directory instead.

>You can also use the Python interpreter of your choice (like python2.7).

```shell    

virtualenv -p /usr/bin/python2.7 venv
```

> or change the interpreter globally with an env variable in ~/.bashrc:

```shell

export VIRTUALENVWRAPPER_PYTHON=/usr/bin/python2.7
``` 

### Virtual Environment Wrapper

>virtualenvwrapper provides a set of commands which makes working with virtual environments much more pleasant. It also places all your virtual environments in one place.

```shell

$ pip install virtualenvwrapper
$ export WORKON_HOME=~/Envs
$ source /usr/local/bin/virtualenvwrapper.sh
```

> Basic usage :

```shell

mkvirtualenv project_folder
workon project_folder
mkproject project_folder
```
