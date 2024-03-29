import { siteConfig } from "@/config/site";
import { Icons } from "@/components/icons";
import { cn } from "@/lib/utils";
import { Link, useMatchRoute, useParams } from "@tanstack/react-router";

export function WorkspaceNav() {
  const matchRoute = useMatchRoute();

  const inWorkspace = matchRoute({
    to: "/workspace/$namespace",
    fuzzy: true,
  });

  const params = useParams({ from: "/_protected/workspace/$namespace" });
  const scope = params.namespace;

  return (
    <div className="mr-4 hidden md:flex">
      <Link to="/workspace" className="mr-6 flex items-center space-x-2">
        <Icons.logo className="h-6 w-6" />
        <span className="hidden font-bold sm:inline-block">
          {siteConfig.name}
        </span>
      </Link>

      {inWorkspace && scope && (
        <nav className="flex items-center gap-6 text-sm">
          <Link
            to={"/workspace/$namespace"}
            params={{ namespace: scope }}
            className={cn(
              "transition-colors hover:text-foreground/80",
              matchRoute({ to: "/workspace/$namespace" })
                ? "text-foreground"
                : "text-foreground/60",
            )}
          >
            Dashboard
          </Link>
          <Link
            to={"/workspace/$namespace/transactions"}
            params={{ namespace: scope }}
            className={cn(
              "transition-colors hover:text-foreground/80",
              matchRoute({ to: "/workspace/$namespace/transactions" })
                ? "text-foreground"
                : "text-foreground/60",
            )}
          >
            Transactions
          </Link>
          <Link
            to={"/workspace/$namespace/accounts"}
            params={{ namespace: scope }}
            className={cn(
              "transition-colors hover:text-foreground/80",
              matchRoute({ to: "/workspace/$namespace/accounts" })
                ? "text-foreground"
                : "text-foreground/60",
            )}
          >
            Accounts
          </Link>

          <Link
            to={"/workspace/$namespace/categories"}
            params={{ namespace: scope }}
            className={cn(
              "transition-colors hover:text-foreground/80",
              matchRoute({ to: "/workspace/$namespace/categories" })
                ? "text-foreground"
                : "text-foreground/60",
            )}
          >
            Categories
          </Link>

          {/* <LinkComponent */}
          {/*   to={"/workspace/$namespace/categories"} */}
          {/*   params={{ namespace: scope }} */}
          {/*   highlight={matchRoute({ to: "/workspace/$namespace/transactions" })} */}
          {/*   name={"Categories"} */}
          {/* /> */}
        </nav>
      )}
    </div>
  );
}

// function LinkComponent(
//   props: LinkProps & { highlight: boolean; name: string },
// ) {
//   return (
//     <Link
//       className={cn(
//         "transition-colors hover:text-foreground/80",
//         props.highlight ? "text-foreground" : "text-foreground/60",
//       )}
//       {...props}
//     >
//       {props.name}
//     </Link>
//   );
// }
