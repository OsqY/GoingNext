import { useAuth } from "@/hooks/useAuth";

export function withAuth<T extends object>(WrappedComponent: ComponentType<T>) {
  return function WithAuth(props: T) {
    const { user, loading } = useAuth();

    if (loading) {
      return <div>Loading...</div>;
    }

    if (!user) {
      return null;
    }

    return <WrappedComponent {...props} user={user} />;
  };
}
