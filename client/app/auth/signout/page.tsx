'use client'

import { useRouter } from "next/navigation"
import { useCallback, useEffect } from "react"
import { fetchApi } from "@/lib/apiWrapper";

const SignOut = () => {
  const router = useRouter();

  const signOutAction = useCallback(async () => {
    try {
      await fetchApi('/api/users/signout');
      router.push('/');
    } catch (error) {
      console.error('Failed', error);
    }
  }, [router]);

  useEffect(() => {
    signOutAction();
  }, [signOutAction]);

  return <div>Signing out...</div>;
};
export default SignOut