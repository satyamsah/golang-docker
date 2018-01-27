
## Part-1 single-file-wordcount
1. `cd single-file-wordcount`

2. `sudo docker build -t single-file-wc .`

3. `sudo docker run --name single-file-wc-ctr -p 9080:9080 single-file-wc`

This is running in non-deamon mode for troubleshooting and log analysis.

4. **Accesing the output**: I exposed the outputfile as a web service. So, Open browser and type :

   `http://localhost:9080/`

5. Optional: The output will dynamically generated in the docker container in the `output` directory.Enter into the container using following command:
   `docker exec -it /bin/bash single-file-wc-ctr` and `cd /output` to see the output file

6. Pre-saved sample outfile "single-file-sample-output" for reference.

## Part-2 multiple_count_wordcount using goroutine

1. `cd ..` to come to main directory

2. `cd multiple_count_wordcount`

2. `sudo docker build -t multiple-files-wc .`

3. `sudo docker run --name multiple-files-wc-ctr -p 9081:9081 multiple-files-wc`

This is running in non-deamon mode for troubleshooting and log analysis.

5.  **Accesing the output**: I exposed the output as a web service. So, Open browser and type:

     `http://localhost:9081/multiplefileswordcount`

4. Optional: The output will be dynamically generated in the docker container in the `output` directory. Enter into the container using following command:
   `docker exec -it /bin/bash multiple-files-wc-ctr` and `cd /output` to see the output file

6. Pre-saved sample outfile "multiplefiles-sample-output" for refrence.
