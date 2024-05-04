import invariant from "tiny-invariant";
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
import {
  Category,
  UpdateCategoryByIdRequest,
} from "@/gen/proto/fijoy/v1/category_pb";
import { ReactNode, useEffect, useRef, useState } from "react";
import { cn } from "@/lib/utils";
import {
  callUnaryMethod,
  createConnectQueryKey,
  useTransport,
} from "@connectrpc/connect-query";
import {
  deleteCategoryById,
  getCategories,
  updateCategoryById,
} from "@/gen/proto/fijoy/v1/category-CategoryService_connectquery";
import { toast } from "sonner";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { useMutation as useConnectMutation } from "@connectrpc/connect-query";
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
  const transport = useTransport();

  const updatePosition = useMutation({
    mutationFn: async (updateCategory: UpdateCategoryByIdRequest) => {
      callUnaryMethod(updateCategoryById, updateCategory, {
        transport,
        callOptions: {
          headers: getWorkspaceHeader(workspace.id),
        },
      });
    },
    onMutate: async (updateCategory) => {
      // Cancel any outgoing refetches
      // (so they don't overwrite our optimistic update)
      const categoriesQueryKey = createConnectQueryKey(getCategories);
      await queryClient.cancelQueries({
        queryKey: categoriesQueryKey,
      });

      const optimisticCategories = [...categories];

      const x = categories.findIndex((e) => e.id === updateCategory.id);

      const [element] = optimisticCategories.splice(x, 1);

      let y;

      if (updateCategory.beforePoition) {
        y =
          categories.findIndex(
            (e) => e.position === updateCategory.beforePoition,
          ) - 1;
      } else {
        y = categories.findIndex(
          (e) => e.position === updateCategory.afterPoition,
        );
      }

      optimisticCategories.splice(y, 0, element);

      queryClient.setQueryData(categoriesQueryKey, optimisticCategories);

      // Return a context object with the snapshotted value
      return { categories };
    },
    // If the mutation fails,
    // use the context returned from onMutate to roll back
    onError: (err, _, context) => {
      console.error(err);
      queryClient.setQueryData(
        createConnectQueryKey(getCategories),
        context?.categories,
      );
    },
    // Always refetch after error or success:
    onSettled: () => {
      queryClient.invalidateQueries({
        queryKey: createConnectQueryKey(getCategories),
      });
    },
  });

  function handlePositionUpdate(
    id: string,
    beforePosition?: string,
    afterPosition?: string,
  ) {
    const req = new UpdateCategoryByIdRequest({
      id,
      beforePoition: beforePosition ?? "",
      afterPoition: afterPosition ?? "",
    });

    updatePosition.mutateAsync(req);
  }

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

function CategoryCardHolder({ children }: { children: ReactNode }) {
  const ref = useRef(null);
  const [isDraggedOver, setIsDraggedOver] = useState(false);

  useEffect(() => {
    const el = ref.current;
    invariant(el);

    return dropTargetForElements({
      element: el,
      onDragEnter: () => setIsDraggedOver(true),

      getData: () => {
        return {
          location: "holder",
        };
      },
      onDrop: ({ location }) => {
        console.log(location);

        setIsDraggedOver(false);
      },
      onDragLeave: () => setIsDraggedOver(false),
    });
  }, []);

  return (
    <div className={cn("", isDraggedOver && "bg-cyan-300")} ref={ref}>
      {children}
    </div>
  );
}

function CategoryCard({ category }: { category: Category }) {
  const ref = useRef(null);
  const [dragging, setDragging] = useState<boolean>(false); // NEW

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
          location: "source",
        };
      },
      onDragStart: () => setDragging(true),
      onDrop: ({ source, location }) => {
        setDragging(false);
        console.log(source, location);
      },
    });
  }, []);

  return (
    <CategoryCardHolder>
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
  const deleteCategory = useConnectMutation(deleteCategoryById, {
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
