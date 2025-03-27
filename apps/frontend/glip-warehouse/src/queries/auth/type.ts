export type RegisterData = {
    email: string;
    password: string;
}

export type RegisterResponse = {
    message: string;
    data: {
        user: {
            id: string;
            email: string;
        }
    }
}

export type LoginData = {
    email: string;
    password: string;
}

export type LoginResponse = {
    message: string;
    data: {
        user: {
            id: string;
            email: string;
        },
        accessToken: string;
    }
}

export type User = {
    id: string;
    email: string;
}