const backend_url  = import.meta.env.VITE_BACKEND_URL;

export async function register(file: File, username: string, password: string): Promise<string> {
    try {
        const uploadedFile = new FormData();
        uploadedFile.append('file', file);
        uploadedFile.append('username', username);
        uploadedFile.append('password', password);


        const response = await fetch(`${backend_url}/register`, {
            method: 'POST',
            body: uploadedFile
        });

        const data = await response.json();
        document.cookie = `jwt=${data.token}; max-age=${data.expiresIn}; path="/";`;

    } catch (error) {
        return error;
    }
}

export async function login(username: string, password: string): Promise<string> {
    try {
        const response = await fetch(`${backend_url}/login`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                username: username,
                password: password
            })
        });

        if (!response.ok) {
            throw new Error('Login failed');
        }

        const data = await response.json();

        document.cookie = `jwt=${data.token}; max-age=${data.expiresIn}; path="/";`;

    } catch (error) {
        return error;
    }
}


/*** Returned error will be null if auth is correct*/
export async function checkAuth(): Promise<string> {
    try {
        const jwtCookie: string = document.cookie.split(';').find(cookie => cookie.trim().startsWith('jwt='));
        const jwtToken: string = jwtCookie ? jwtCookie.split('=')[1] : '';

        const response = await fetch(`${backend_url}/auth`, {
            method: 'GET',
            credentials: 'include',
            headers: {
                'Authorization': `Bearer ${jwtToken}`
            }
        });
        const data = await response.json();

    } catch (error) {
        return error;
    }
}