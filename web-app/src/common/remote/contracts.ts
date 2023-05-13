export declare type NotesApiUrls = {
  signUpUrl: string;
  loginUrl: string;
};

const urls: NotesApiUrls = {
  signUpUrl: "/signup",
  loginUrl: "/login",
};


export declare type ApiResponse<T> = {
    data: T
    error?: string
} 


export { urls };
