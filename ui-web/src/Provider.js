import React, { useState } from "react";
import UserContext from "./UserContext";


const Provider = ({children}) => {
  const [userDetails, setUserDetails] = useState({});
  return (
    <UserContext.Provider value={{userDetails, setUserDetails}}>
      {children}
    </UserContext.Provider>
  )
};

export default Provider;
