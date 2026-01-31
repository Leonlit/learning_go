import { Navigate } from "react-router-dom";
import { ReactNode } from "react";
import useAuth from "../components/auth/useAuth";

interface ProtectedRoutesProps {
    children: ReactNode;
}

const ProtectedRoute = ({ children }: ProtectedRoutesProps ) => {
    const { authenticated, loading } = useAuth();

    if (loading) return <div>Loading...</div>;
    if (!authenticated) return <Navigate to="/login" />;
    return children;
};

export default ProtectedRoute;