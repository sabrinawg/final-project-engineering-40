### CEK API

- `http://localhost:8080/cek` => (GET) cek

### AUTH ADMIN/UNIVERSITAS

- `http://localhost:8080/signup` => (POST) signup
- `http://localhost:8080/signin` => (POST) signin

### ADMIN/UNIVERSITAS

- `http://localhost:8080/admin/user` => (GET) CurrentUser
- `http://localhost:8080/admin/user` => (PUT) Update Admin/Univ
- `http://localhost:8080/admin/user/mhs` => (POST) Create mahasiswa
- `http://localhost:8080/admin/user/mhs/:id` => (DELETE) mahasiswa
- `http://localhost:8080/admin/user/fakultas` => (POST) Create Fakultas
- `http://localhost:8080/admin/user/fakultas/:id` => (UPDATE) Fakultas
- `http://localhost:8080/admin/user/fakultas/:id` => (DELETE) Fakultas
- `http://localhost:8080/admin/user/prodi` => (POST) Create Prodi
- `http://localhost:8080/admin/user/prodi/:id` => (UPDATE) Prodi
- `http://localhost:8080/admin/user/prodi/:id` => (DELETE) Prodi
- `http://localhost:8080/admin/user/post` => (POST) Create Post

### USER

- `http://localhost:8080/user` => (GET) all user

### MAHASISWA

- `http://localhost:8080/mhs/signin` => (POST) mahasiswa signin
- `http://localhost:8080/mhs/post` => (POST) create post

### Get fakultas

- `http://localhost:8080/fakultas` => (GET) all fakultas
- `http://localhost:8080/fakultas/:id` => (GET) single fakultas by id
- `http://localhost:8080/fakultas/name/:name_fakultas` => (GET) single fakultas by name

### Get Prodi

- `http://localhost:8080/prodi` => (GET) all prodi
- `http://localhost:8080/prodi/:id` => (GET) single prodi by id
- `http://localhost:8080/prodi/name/:name_prodi` => (GET) single prodi by name

### POST_UNIV

- `http://localhost:8080/post` => (GET) all post universitas

### POST_MHS

- `http://localhost:8080/postmhs` => (GET) all post mahasiswa
