import { useEffect, useState } from "react";

const useAuth = () => {
    const [loading, setLoading] = useState(true);
    const [authenticated, setAuthenticated] = useState(false);

    useEffect(() => {
        const check = async () => {
            try {
                const res = await fetch("/auth/me", {
                    credentials: "include", // this sends the cookie
                });
                setAuthenticated(res.ok);
            } catch {
                setAuthenticated(false);
            } finally {
                setLoading(false);
            }
        };

        check();
    }, []);

    return { authenticated, loading };
};

export default useAuth