## assigment
A quick 3-hr effort to develop a web server to store the user's upload video files and provide the basic CRUD rest APIs to create/query/delete the uploaded files

## limitation
If I have more time, I would have:
- added the unit tests for each of the handler function
- added the unit tests for each database APIs
- took a simple assumption on reporting the video length as the time taken to upload the file, which may not be correct
- took a simple approach to treat the video file as a small file which can fit into the memory. Normally it needs to be uploaded in multi-part form
- did not consider the time in UTC for simplification
