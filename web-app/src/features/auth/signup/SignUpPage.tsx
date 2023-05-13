import React, { useCallback } from "react";
import useViewModel from "./viewmodel";

const SignUpPage = () => {
  const { name, email, password, signUp } = useViewModel();

  const handleSubmit = useCallback(() => {
    signUp()
  }, [signUp]);

  return (
    <div>
      <input
        type="text"
        value={name.value}
        onChange={(e) => name.setter(e.target.value)}
      />
      <br></br>
      <input
        type="email"
        value={email.value}
        onChange={(e) => email.setter(e.target.value)}
      />
      <br></br>
      <input
        type="password"
        value={password.value}
        onChange={(e) => password.setter(e.target.value)}
      />
      <br></br>
      <button onClick={(e) => handleSubmit()}>Sign up</button>
    </div>
  );
};


export default SignUpPage;