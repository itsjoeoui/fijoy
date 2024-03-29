import { siteConfig } from "@/config/site";
import { cn } from "@/lib/utils";
import { HTMLAttributes, forwardRef } from "react";
import { Icons } from "./icons";
import { Link, useMatchRoute, useParams } from "@tanstack/react-router";
import { ArrowLeftRight, CreditCard, Home, Settings } from "lucide-react";
// import { ModeToggle } from "./mode-toggle";
// import { Button, buttonVariants } from "./ui/button";
import { Button, buttonVariants } from "./ui/button";
// import { useAuth } from "@/auth";
import { Badge } from "./ui/badge";
import { useAuth } from "@/auth";
import WorkspaceSwitcher from "./workspace-switcher";
import UserButton from "./user-button";
import { ModeToggle } from "./mode-toggle";

const PrivateSidebar = forwardRef<
  HTMLDivElement,
  HTMLAttributes<HTMLDivElement>
>(({ className, ...props }, ref) => {
  const matchRoute = useMatchRoute();

  const inWorkspace = matchRoute({
    to: "/workspace/$namespace",
    fuzzy: true,
  });

  const params = useParams({ from: "/_protected/workspace/$namespace" });
  const scope = params.namespace;
  const { user } = useAuth();

  if (!user) {
    return null;
  }

  return (
    <div
      ref={ref}
      className={cn("flex flex-col border p-4", className)}
      {...props}
    >
      <div className="flex items-center">
        <Link to="/" className="mr-6 flex items-center space-x-2">
          <Icons.logo className="h-6 w-6" />
          <span className="text-xl font-bold">{siteConfig.name}</span>
          <Badge>Pro</Badge>
        </Link>
        <div className="grow"></div>
        <UserButton />
      </div>

      <div className="py-2"></div>

      {inWorkspace && scope && (
        <>
          <Button variant="outline">New Transaction</Button>

          <div className="py-2"></div>

          <nav className="flex flex-col gap-2 text-sm font-bold">
            <Link
              to={"/workspace/$namespace"}
              params={{ namespace: scope }}
              className={cn(
                "flex items-center gap-2 rounded-md p-2 transition-colors hover:bg-background hover:text-foreground",
                matchRoute({ to: "/workspace/$namespace/" })
                  ? "text-primary hover:text-primary"
                  : "text-foreground/60",
              )}
            >
              <Home size="20" />
              Home
            </Link>
            <Link
              to={"/workspace/$namespace/transactions"}
              params={{ namespace: scope }}
              className={cn(
                "flex items-center gap-2 rounded-md p-2 transition-colors hover:bg-background hover:text-foreground",
                matchRoute({ to: "/workspace/$namespace/transactions/" })
                  ? "text-primary hover:text-primary"
                  : "text-foreground/60",
              )}
            >
              <ArrowLeftRight size={20} />
              Transactions
            </Link>
            <Link
              to={"/workspace/$namespace/accounts"}
              params={{ namespace: scope }}
              className={cn(
                "flex items-center gap-2 rounded-md p-2 transition-colors hover:bg-background hover:text-foreground",
                matchRoute({ to: "/workspace/$namespace/accounts/" })
                  ? "text-primary hover:text-primary"
                  : "text-foreground/60",
              )}
            >
              <CreditCard size="20" />
              Accounts
            </Link>

            <Link
              to={"/workspace/$namespace/settings"}
              params={{ namespace: scope }}
              className={cn(
                "flex items-center gap-2 rounded-md p-2 transition-colors hover:bg-background hover:text-foreground",
                matchRoute({ to: "/workspace/$namespace/settings/" })
                  ? "text-primary hover:text-primary"
                  : "text-foreground/60",
              )}
            >
              <Settings size="20" />
              Settings
            </Link>

            {/* <Link */}
            {/*   to={"/workspace/$namespace/categories"} */}
            {/*   params={{ namespace: scope }} */}
            {/*   className={cn( */}
            {/*     "transition-colors hover:text-foreground/80", */}
            {/*     matchRoute({ to: "/workspace/$namespace/categories" }) */}
            {/*       ? "text-foreground" */}
            {/*       : "text-foreground/60", */}
            {/*   )} */}
            {/* > */}
            {/*   Categories */}
            {/* </Link> */}

            {/* <LinkComponent */}
            {/*   to={"/workspace/$namespace/categories"} */}
            {/*   params={{ namespace: scope }} */}
            {/*   highlight={matchRoute({ to: "/workspace/$namespace/transactions" })} */}
            {/*   name={"Categories"} */}
            {/* /> */}
          </nav>
        </>
      )}

      <div className="grow"></div>
      <div className="flex">
        <a href={siteConfig.links.github} target="_blank" rel="noreferrer">
          <div
            className={cn(buttonVariants({ variant: "ghost", size: "icon" }))}
          >
            <Icons.gitHub className="h-[1.2rem] w-[1.2rem] transition-all hover:scale-125" />
            <span className="sr-only">GitHub</span>
          </div>
        </a>
        <a href={siteConfig.links.discord} target="_blank" rel="noreferrer">
          <div
            className={cn(buttonVariants({ variant: "ghost", size: "icon" }))}
          >
            <Icons.discord className="h-[1.2rem] w-[1.2rem] transition-all hover:scale-125" />
            <span className="sr-only">Discord</span>
          </div>
        </a>

        <div className="grow"></div>

        <ModeToggle />
      </div>
      <div className="py-1"></div>

      <WorkspaceSwitcher />
    </div>
  );
});

export default PrivateSidebar;
