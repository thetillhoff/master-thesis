FROM gitpod/workspace-base:latest

RUN sudo sh -c "$(curl --location https://taskfile.dev/install.sh)" -- -b /usr/local/bin

RUN apt-get install python-pip latexmk -y

RUN apt-get install texlive-xetex texlive texlive-latex-extra texlive-generic-extra texlive-lang-german texlive-bibtex-extra biber -y
# Following line is not used, as it downloads EVERY language and therefore requires several hundred megabytes
#RUN apt-get install texlive-full

RUN pip install pygments

RUN mkdir /temp

ENV textype=pdf
ENV filename=main.tex

WORKDIR /temp

# ? eventually latexmk command has to executed twice for completing properly
CMD ["/bin/bash","-c","cp -dr /tex/* /temp/ && rm /temp/*.pdf || true && latexmk -shell-escape -$textype $filename && cp /temp/main.pdf /tex/ || cp /temp/main.log /tex/"]
# The following line would output only the last part of the compiling output. This may be better for beginners but hides progress and hides earlier warnings.
#CMD ["/bin/bash","-c","cp -dr /tex/* /temp/ && rm /temp/*.pdf && latexmk -shell-escape -$textype $filename | tail | sed '/^============$/ q' && cp /temp/main.pdf /tex/ || cp /temp/main.log /tex/"]