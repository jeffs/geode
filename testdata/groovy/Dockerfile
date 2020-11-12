FROM {{.Base}}

# Restore man pages, etc. stripped from Ubuntu to minimize default image size.
RUN yes | unminimize

# Check package availability.  To rerun, pass --no-cache to docker image build.
RUN apt-get update -qq

# Set locale.
RUN DEBIAN_FRONTEND=noninteractive apt-get install -qqy locales \
 && locale-gen {{.Locale}} \
 && dpkg-reconfigure --frontend noninteractive locales \
 && echo LC_ALL={{.Locale}} >>/etc/default/locale
ENV LANG={{.Locale}}

# Set time zone.
RUN DEBIAN_FRONTEND=noninteractive apt-get install -qqy tzdata \
 && ln -snf /usr/share/zoneinfo/{{.TimeZone}} /etc/localtime \
 && echo {{.TimeZone}} > /etc/timezone \
 && dpkg-reconfigure --frontend noninteractive tzdata
ENV TZ={{.TimeZone}}

# Install command-line tools.  TODO: Configure per TOML.
#
# programming languages:
#       c               # C and C++ toolchain and standard library
#       (build-essential)
#
#       python3         # Python 3 with pip, venv, and headers
#       (python3-pip python3-venv)
#
# network:
#       curl            # HTTP utility
#       netcat          # TCP/IP Swiss army knife
#       (netcat-traditional)
#
# basic commands:
#       git             # version control system
#       man             # manual page viewer
#
# file types:
#       jq              # parses, queries, transforms, and formats JSON
#       markdown        # converts Markdown files to HTML
#       zip             # creates or expands .zip compressed archives
#
# enhanced commands:
#       bat             # alternative to cat
#       exa             # alternative to ls and tree
#
#       neovim          # fork of vim
#
#       fd (fd-find)    # alternative to find
#       ripgrep         # alternative to grep
#
#       sl              # steam locomotive when you mistype ls
#       tmux            # terminal multiplexer
#       zsh             # alternative to sh
#
RUN DEBIAN_FRONTEND=noninteractive apt-get install -qqy \
    build-essential python3-venv \
    bat exa fd-find git neovim ripgrep sl tmux zsh \
    curl netcat-traditional net-tools sudo \
    jq man markdown zip

# Install Geode guest scripts.
# COPY sbin /sbin

{{- if .User}}
# Create myself, and give myself super powers.
#
# Enabling sudo coredumps works around a container bug that makes sudo unhappy:
# https://github.com/sudo-project/sudo/issues/42
RUN useradd -ms /usr/bin/zsh {{.User}} \
 && echo '{{.User}} ALL=(ALL) NOPASSWD:ALL' >> /etc/sudoers \
 && echo "Set disable_coredump false" >> /etc/sudo.conf

# Make myself the default user.
USER {{.User}}
{{- end}}

# TODO Make shell configurable.
CMD ["/usr/bin/zsh"]
