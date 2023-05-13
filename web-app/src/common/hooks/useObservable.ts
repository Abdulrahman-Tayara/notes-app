import { useState } from "react";

export declare type Observable<S> = {
  value: S
  setter: React.Dispatch<React.SetStateAction<S>>
}

function useObservable<S>(defaultValue: S): Observable<S> {
  const [value, setter] = useState<S>(defaultValue);

  return {value, setter};
}

export { useObservable };
