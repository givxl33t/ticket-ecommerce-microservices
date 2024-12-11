"use client";

import { useContext, createContext, useState, useEffect } from "react";
import axios from "axios";

interface IAuthContext {
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  currentUser: any;
  refreshUser: () => void;
}

export const AuthContext = createContext<IAuthContext>(null!);

export const AuthProvider = ({ children }: { children: React.ReactNode }) => {
  const [currentUser, setCurrentUser] = useState(null);

  const refreshUser = async () => {
    try {
      const response = await axios.get("/api/users/currentuser");
      const data = response.data;
      setCurrentUser(data.currentUser);
    } catch (error) {
      console.error(error);
    }
  };

  useEffect(() => {
    refreshUser();
  }, []);

  return (
    <AuthContext.Provider value={{ currentUser, refreshUser }}>{children}</AuthContext.Provider>
  );
};

export const useAuth = () => useContext(AuthContext);
