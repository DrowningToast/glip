"use client";

import { zodResolver } from "@hookform/resolvers/zod";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { Loader2, Plus } from "lucide-react";
import { useState } from "react";
import { useForm } from "react-hook-form";
import { z } from "zod";
import { registryApi } from "../api/registry";

import { toast } from "sonner";
import { Button } from "./ui/button";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "./ui/dialog";
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "./ui/form";
import { Input } from "./ui/input";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "./ui/select";

// Define the form schema with Zod
const warehouseFormSchema = z.object({
  warehouse_id: z.string().min(1, "Warehouse ID is required"),
  api_key: z.string().min(1, "API Key is required"),
  name: z.string().min(1, "Name is required"),
  status: z.enum(["ACTIVE", "INACTIVE"]),
});

type WarehouseFormValues = z.infer<typeof warehouseFormSchema>;

// Default values for the form
const defaultValues: WarehouseFormValues = {
  warehouse_id: "",
  api_key: "",
  name: "",
  status: "ACTIVE",
};

export default function WarehouseDialog() {
  const [open, setOpen] = useState(false);
  const [isSubmitting, setIsSubmitting] = useState(false);
  const queryClient = useQueryClient();
  const mutation = useMutation({
    mutationFn: (data: WarehouseFormValues) =>
      registryApi.warehouseConnection.createConnection({
        headers: {
          authorization: `${import.meta.env.VITE_REGISTRY_SECRET_KEY || ""}`,
          authtype: "ADMIN",
        },
        body: {
          warehouse_connection: {
            warehouse_id: data.warehouse_id,
            api_key: data.api_key,
            name: data.name,
            status: data.status,
          },
        },
      }),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["warehouses"] });
    },
  });

  // Initialize the form
  const form = useForm<WarehouseFormValues>({
    resolver: zodResolver(warehouseFormSchema),
    defaultValues,
  });

  // Handle form submission
  const onSubmit = async (data: WarehouseFormValues) => {
    setIsSubmitting(true);
    mutation
      .mutateAsync(data)
      .then((res) => {
        console.log("Warehouse created:", res);

        // Show success message
        toast.success(`Successfully created warehouse: ${data.name}`);

        // Reset form and close dialog
        form.reset(defaultValues);
        setOpen(false);
        setIsSubmitting(false);
      })
      .catch((err) => {
        console.error("Error creating warehouse:", err);
        toast.error("Failed to create warehouse. Please try again.");
        setIsSubmitting(false);
      });
  };

  return (
    <Dialog open={open} onOpenChange={setOpen}>
      <DialogTrigger asChild>
        <Button variant="outline">
          <Plus className="mr-2 h-4 w-4" />
          Add Warehouse
        </Button>
      </DialogTrigger>
      <DialogContent className="sm:max-w-[425px]">
        <DialogHeader>
          <DialogTitle>Create Warehouse</DialogTitle>
          <DialogDescription>
            Add a new warehouse to your inventory management system.
          </DialogDescription>
        </DialogHeader>
        <Form {...form}>
          <form
            onSubmit={form.handleSubmit(onSubmit)}
            className="space-y-4 py-4"
          >
            <FormField
              control={form.control}
              name="warehouse_id"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Warehouse ID</FormLabel>
                  <FormControl>
                    <Input placeholder="USA1" {...field} />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />

            <FormField
              control={form.control}
              name="api_key"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>API Key</FormLabel>
                  <FormControl>
                    <Input placeholder="Enter API Key" {...field} />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />

            <FormField
              control={form.control}
              name="name"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Name</FormLabel>
                  <FormControl>
                    <Input placeholder="USA Warehouse 1" {...field} />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />

            <FormField
              control={form.control}
              name="status"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Status</FormLabel>
                  <Select
                    onValueChange={field.onChange}
                    defaultValue={field.value}
                  >
                    <FormControl>
                      <SelectTrigger>
                        <SelectValue placeholder="Select status" />
                      </SelectTrigger>
                    </FormControl>
                    <SelectContent>
                      <SelectItem value="ACTIVE">ACTIVE</SelectItem>
                      <SelectItem value="INACTIVE">INACTIVE</SelectItem>
                    </SelectContent>
                  </Select>
                  <FormMessage />
                </FormItem>
              )}
            />

            <DialogFooter>
              <Button
                type="button"
                variant="outline"
                onClick={() => setOpen(false)}
              >
                Cancel
              </Button>
              <Button type="submit" disabled={isSubmitting}>
                {isSubmitting ? (
                  <>
                    <Loader2 className="mr-2 h-4 w-4 animate-spin" />
                    Creating...
                  </>
                ) : (
                  "Create Warehouse"
                )}
              </Button>
            </DialogFooter>
          </form>
        </Form>
      </DialogContent>
    </Dialog>
  );
}
