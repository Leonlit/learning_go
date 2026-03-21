import { useEffect, useState } from "react";

const TeamMemberWidget = () => {
    const [loading, setLoading] = useState(true);
    const [authenticated, setAuthenticated] = useState(false);

    useEffect(() => {
        const check = async () => {
            try {
                const res = await fetch("http://localhost:8080/api/users/team-member", {
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

    return (
        <div>
            
        </div>
    );
};

export default TeamMemberWidget