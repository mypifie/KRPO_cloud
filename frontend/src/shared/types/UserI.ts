export interface UserI {
    email: string | null;
    username: string;
    isAdmin: boolean;
    isBlocked: boolean;
}

export interface RegistrationDataI {
  username: string;
  email: string;
  password: string;
}

export interface AuthDataI {
  email: string;
  password: string;
}
