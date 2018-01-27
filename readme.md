
## Part-1 single-file-wordcount
1. `cd single-file-wordcount`
2. `sudo docker build -t single-file-wc .`
3. `sudo docker run --name single-file-wc-ctr -p 9080:9080 single-file-wc`. This is running in non-deamon mode for troubleshooting and log analysis.
4. The output will saved in the docker container in the `output` directory.
   `docker exec -it /bin/bash single-file-wc-ctr`
4. I exposed the output as a web service. So, Open browser and type :
   `http://localhost:9080/`
6. Sample output in csv is named as "sinlge-file-sample-output". Pre-saved by me

## Part-2 multiple_count_wordcount using goroutine
1. `cd ..` to come to main directory
2. `sudo docker build -t multiple-files-wc .`
3. `sudo docker run --name multiple-files-wc-ctr -p 9081:9081 multiple-files-wc`. This is running in non-deamon mode for troubleshooting and log analysis.
4. The output will saved in the docker container in the `output` directory.
   `docker exec -it /bin/bash multiple-files-wc-ctr`
5.  I exposed the output as a web service. So, Open browser and type:
     `http://localhost:9081/multiplefileswordcount`
6. Sample output in csv is named as "multiplefiles-sample-output".Pre-saved by me
