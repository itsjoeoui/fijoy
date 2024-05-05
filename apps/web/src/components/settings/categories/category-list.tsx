import invariant from "tiny-invariant";
import {
  attachClosestEdge,
  extractClosestEdge,
  Edge,
} from "@atlaskit/pragmatic-drag-and-drop-hitbox/closest-edge";
import {
  AlertDialog,
  AlertDialogAction,
  AlertDialogCancel,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
  AlertDialogTrigger,
} from "@/components/ui/alert-dialog";
import {
  draggable,
  dropTargetForElements,
  monitorForElements,
} from "@atlaskit/pragmatic-drag-and-drop/element/adapter";
import {
  AccordionContent,
  AccordionItem,
  AccordionTrigger,
} from "@/components/ui/accordion";
import { GripVertical, Plus, Receipt, Trash } from "lucide-react";
import { Card, CardContent } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import _ from "lodash";
import { TransactionTypeEnum } from "@/types/transaction";
import { Category } from "@/gen/proto/fijoy/v1/category_pb";
import { ReactNode, useEffect, useRef, useState } from "react";
import { cn } from "@/lib/utils";
import { createConnectQueryKey, useMutation } from "@connectrpc/connect-query";
import {
  deleteCategoryById,
  getCategories,
  updateCategoryById,
} from "@/gen/proto/fijoy/v1/category-CategoryService_connectquery";
import { toast } from "sonner";
import { useQueryClient } from "@tanstack/react-query";
import { getWorkspaceHeader } from "@/lib/headers";
import { useWorkspace } from "@/hooks/use-workspace";

type CategoryListProps = {
  categories: Category[];
  categoryType: TransactionTypeEnum;
  onAddNewCategory: (categoryType: TransactionTypeEnum) => void;
};
export function CategoryList({
  categories,
  categoryType,
  onAddNewCategory,
}: CategoryListProps) {
  const { workspace } = useWorkspace();

  const queryClient = useQueryClient();

  const updatePosition = useMutation(updateCategoryById, {
    callOptions: {
      headers: getWorkspaceHeader(workspace.id),
    },
    onSuccess: () => {
      queryClient.invalidateQueries({
        queryKey: createConnectQueryKey(getCategories),
      });
    },
  });

  function getBeforeAfterPosition(
    referencePosition: string,
    location: "top" | "bottom",
  ) {
    const referenceIndex = categories.findIndex(
      (c) => c.position === referencePosition,
    );
    if (referenceIndex === -1) {
      throw new Error("Reference category not found");
    }

    if (location === "top") {
      const targetIndex = referenceIndex - 1;
      if (targetIndex < 0) {
        return {
          beforePosition: referencePosition,
          afterPosition: "",
        };
      }

      return {
        beforePosition: referencePosition,
        afterPosition: categories[targetIndex].position,
      };
    }
    if (location === "bottom") {
      const targetIndex = referenceIndex + 1;
      if (targetIndex >= categories.length) {
        return {
          beforePosition: "",
          afterPosition: referencePosition,
        };
      }

      return {
        beforePosition: categories[targetIndex].position,
        afterPosition: referencePosition,
      };
    }
  }

  // this makes sure that the mutation only fires once in strict mode
  useEffect(() => {
    return monitorForElements({
      onDrop({ source, location }) {
        const destination = location.current.dropTargets[0];
        if (!destination) {
          // if dropped outside of any drop targets
          return;
        }
        const destinationLocation = destination.data.location;
        const sourceLocation = source.data.location;
        console.log("---------");
        console.log("source.data", source.data);
        console.log("destination.data", destination.data);

        console.log(
          "calling with",
          source.data.category.id,
          destination.data.position,
        );

        updatePosition.mutateAsync({
          id: source.data.category.id,
          ...getBeforeAfterPosition(destination.data.position, "top"),
          // afterPosition: destination.data.position ?? "",
          // beforePosition: "",
        });

        // if (
        //   // type guarding
        //   !isCoord(destinationLocation) ||
        //   !isCoord(sourceLocation) ||
        //   !isPieceType(pieceType)
        // ) {
        //   return;
        // }

        // const piece = pieces.find((p) =>
        //   isEqualCoord(p.location, sourceLocation),
        // );
        // const restOfPieces = pieces.filter((p) => p !== piece);
        //
        // if (
        //   canMove(sourceLocation, destinationLocation, pieceType, pieces) &&
        //   piece !== undefined
        // ) {
        //   // moving the piece!
        //   setPieces([
        //     { type: piece.type, location: destinationLocation },
        //     ...restOfPieces,
        //   ]);
        // }
      },
    });
  }, [categories]);

  return (
    <Card className="md:min-w-80">
      <CardContent className="pb-0">
        <AccordionItem value={categoryType} className="border-none ">
          <AccordionTrigger>
            <Receipt className="h-4 w-4 !transform-none" />
            {_.capitalize(categoryType)}
          </AccordionTrigger>
          <AccordionContent className="space-y-2">
            <Button
              variant="secondary"
              className="w-full gap-1"
              onClick={() => onAddNewCategory(categoryType)}
            >
              <Plus className="h-4 w-4" />
              New
            </Button>
            {categories.map((category) => (
              <CategoryCard category={category} key={category.id} />
            ))}
          </AccordionContent>
        </AccordionItem>
      </CardContent>
    </Card>
  );
}

function CategoryCardHolder({
  position,
  children,
}: {
  position: string;
  children: ReactNode;
}) {
  const ref = useRef(null);
  const [isDraggedOver, setIsDraggedOver] = useState(false);

  useEffect(() => {
    const el = ref.current;
    invariant(el);

    return dropTargetForElements({
      element: el,
      onDragEnter: () => setIsDraggedOver(true),

      getData: ({ input, element }) => {
        // your base data you want to attach to the drop target
        const data = {
          position,
        };
        // this will 'attach' the closest edge to your `data` object
        return attachClosestEdge(data, {
          input,
          element,
          // you can specify what edges you want to allow the user to be closest to
          allowedEdges: ["top", "bottom"],
        });
      },
      onDrop: ({ location, self }) => {
        // console.log(location);
        const closestEdgeOfTarget: Edge | null = extractClosestEdge(self.data);
        // console.log(closestEdgeOfTarget);
        setIsDraggedOver(false);
      },
      onDragLeave: () => setIsDraggedOver(false),
    });
  }, [position]);

  return (
    <div className={cn("", isDraggedOver && "bg-cyan-300")} ref={ref}>
      {children}
    </div>
  );
}

function CategoryCard({ category }: { category: Category }) {
  const ref = useRef(null);

  const [dragging, setDragging] = useState<boolean>(false);

  const dragHandleRef = useRef<HTMLButtonElement>(null);

  useEffect(() => {
    const el = ref.current;
    invariant(el);
    invariant(dragHandleRef.current);

    return draggable({
      element: el,
      dragHandle: dragHandleRef.current,
      getInitialData() {
        return {
          category,
        };
      },
      onDragStart: () => setDragging(true),
      onDrop: ({ source, location }) => {
        setDragging(false);
        // console.log(source, location);
      },
    });
  }, []);

  return (
    <CategoryCardHolder position={category.position}>
      <Card
        className={cn("flex items-center p-2", dragging && "bg-muted")}
        ref={ref}
      >
        <div>
          {category.name}, {category.position}
        </div>
        <div className="grow"></div>

        <DeleteCategoryButton id={category.id} />

        <div className="px-1"></div>

        <Button
          ref={dragHandleRef}
          variant="ghost"
          size="icon"
          className="h-4 w-4 hover:bg-background"
        >
          <GripVertical className="h-4 w-4" />
        </Button>
      </Card>
    </CategoryCardHolder>
  );
}

function DeleteCategoryButton({ id }: { id: string }) {
  const queryClient = useQueryClient();
  const { workspace } = useWorkspace();
  const deleteCategory = useMutation(deleteCategoryById, {
    callOptions: {
      headers: getWorkspaceHeader(workspace.id),
    },
    onSuccess: () => {
      queryClient.invalidateQueries({
        queryKey: createConnectQueryKey(getCategories),
      });
    },
  });

  return (
    <AlertDialog>
      <AlertDialogTrigger asChild>
        <Button
          variant="ghost"
          size="icon"
          className="h-4 w-4 hover:bg-background"
        >
          <Trash className="h-4 w-4" />
        </Button>
      </AlertDialogTrigger>
      <AlertDialogContent>
        <AlertDialogHeader>
          <AlertDialogTitle>Are you absolutely sure?</AlertDialogTitle>
          <AlertDialogDescription>
            This action cannot be undone. All transactions under this category
            will lose its category information.
          </AlertDialogDescription>
        </AlertDialogHeader>
        <AlertDialogFooter>
          <AlertDialogCancel>Cancel</AlertDialogCancel>
          <AlertDialogAction
            onClick={() =>
              toast.promise(deleteCategory.mutateAsync({ id }), {
                success: "Category deleted!",
                error: "Failed to delete category",
              })
            }
          >
            Continue
          </AlertDialogAction>
        </AlertDialogFooter>
      </AlertDialogContent>
    </AlertDialog>
  );
}