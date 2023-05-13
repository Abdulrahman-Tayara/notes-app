import useCommand from "common/hooks/useCommand";
import React, { useState } from "react";
import { SignUpCommand } from "./commands";
import { Observable, useObservable } from "common/hooks/useObservable";

interface ISignUPViewModel {
    email: Observable<string>
    name: Observable<string>
    password: Observable<string>

    signUp(): void
}

export default function useViewModel(): ISignUPViewModel {
    const signUpCommand = useCommand<SignUpCommand>(SignUpCommand);

    const email = useObservable<string>("");
    const name = useObservable<string>("");
    const password = useObservable<string>("");


    const signUp = () => {
        signUpCommand.handle({
            name: name.value,
            password: password.value,
            email: email.value
        }).then(user => {
            console.log(user);
        })
    };

    return {
        email,
        name,
        password,
        signUp
    }
}