import { createFileRoute, redirect, useNavigate } from "@tanstack/react-router";
import { useQuery, useQueryClient } from "@tanstack/react-query";
import { useState } from "react";

import { getUser, logout } from "@/queries/auth/query";
import { getAllShipment } from "@/queries/inventory/query";
import { Button } from "@/components/ui/button";
import {
  ColumnDef,
  flexRender,
  getCoreRowModel,
  useReactTable,
  getPaginationRowModel,
  getSortedRowModel,
  SortingState,
  getFilteredRowModel,
  ColumnFiltersState,
} from "@tanstack/react-table";
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";

import { ArrowUpDown, RotateCw } from "lucide-react";
import { Input } from "@/components/ui/input";
import { cn } from "@/libs/utils";
import { UpdateStatusDialog } from "@/components/UpdateStatusDialog";
import { Inventory } from "@/queries/inventory/type";
import { Badge } from "@/components/ui/badge";

export const Route = createFileRoute("/")({
  component: Index,
  loader: async () => {
    try {
      const user = await getUser();
      if (!user) {
        return redirect({ to: "/login" });
      }
      return null;
    } catch (error) {
      console.error(error);
      return redirect({ to: "/login" });
    }
  },
});

const columns: ColumnDef<Inventory>[] = [
  {
    header: "Shipment ID",
    accessorKey: "shipmentId",
  },
  {
    header: "Status",
    accessorKey: "Status",
    cell: ({ row }) => {
      switch (row.original.Status) {
        case "DELIVERED":
          return <Badge variant="default" className="bg-green-500">Delivered</Badge>;
        case "CANCELLED":
          return <Badge variant="destructive" className="bg-red-500">Cancelled</Badge>;
        case "INCOMING_SHIPMENT":
          return <Badge variant="default" className="bg-blue-500">Incoming Shipment</Badge>;
        case "WAREHOUSE_RECEIVED":
          return <Badge variant="default" className="bg-yellow-500">Warehouse Received</Badge>;
        case "WAREHOUSE_DEPARTED":
          return <Badge variant="default" className="bg-yellow-500">Warehouse Departed</Badge>;
      }
    },
  },
  {
    header: "From Warehouse",
    accessorKey: "FromWarehouseId",
  },
  {
    header: "To Warehouse",
    accessorKey: "ToWarehouseId",
  },
  {
    header: "Route",
    accessorKey: "Route",
  },
  {
    header: "Owner",
    accessorKey: "OwnerId",
  },
  {
    header: "Special Instructions",
    accessorKey: "SpecialInstructions",
    cell: ({ row }) => {
      return <div>{row.original.SpecialInstructions || "-"}</div>;
    },
  },
  {
    header: "Created By",
    accessorKey: "CreatedBy",
  },
  {
    header: ({ column }) => (
      <Button
        variant="ghost"
        onClick={() => column.toggleSorting(column.getIsSorted() === "asc")}
      >
        Delivery Time
        <ArrowUpDown className="ml-2 h-4 w-4" />
      </Button>
    ),
    accessorKey: "DeliveryTime",
    cell: ({ row }) => {
      return (
        <div>
          {row.original.DeliveryTime ? new Date(row.original.DeliveryTime).toLocaleString("th-TH", {
            year: "numeric",
            month: "2-digit",
            day: "2-digit",
            hour: "2-digit",
            minute: "2-digit",
          }) : "-"}
        </div>
      );
    },
  },
  {
    header: "Action",
    accessorKey: "Action",
    cell: ({ row }) => {
      return (
        <div>
          <UpdateStatusDialog
            id={row.original.id}
            shipmentId={row.original.shipmentId}
            status={row.original.Status}
          />
        </div>
      );
    },
  },
];

function Index() {
  const [sorting, setSorting] = useState<SortingState>([]);
  const [filtering, setFiltering] = useState<ColumnFiltersState>([]);
  const [lastRefresh, setLastRefresh] = useState<Date>(new Date());
  const queryClient = useQueryClient();
  const navigate = useNavigate();
  const [isRefreshing, setIsRefreshing] = useState(false);

  const { data } = useQuery({
    queryKey: ["shipments"],
    queryFn: getAllShipment,
  });

  const table = useReactTable({
    data: data || [],
    columns,
    getCoreRowModel: getCoreRowModel(),
    getPaginationRowModel: getPaginationRowModel(),
    getSortedRowModel: getSortedRowModel(),
    onSortingChange: setSorting,
    state: {
      sorting,
      columnFilters: filtering,
    },
    getFilteredRowModel: getFilteredRowModel(),
    onColumnFiltersChange: setFiltering,
  });

  const handleRefresh = async () => {
    setIsRefreshing(true);
    await queryClient.invalidateQueries({ queryKey: ["shipments"] });
    setLastRefresh(new Date());
    setIsRefreshing(false);
  };

  const handleLogout = async () => {
    await logout();
    navigate({ to: "/login" });
  };

  return (
    <div className="max-w-7xl mx-auto py-14">
      <div className="flex justify-between">
        <h1 className="text-2xl font-bold">Inventory Dashboard</h1>
        <Button variant="outline" size="sm" onClick={handleLogout}>Logout</Button>
      </div>
      <div className="mt-4">
        <div className="flex justify-between items-center py-4">
          <div className="flex items-center gap-2">
            <Input
              placeholder="Filter Shipment ID..."
              value={
                (table.getColumn("shipmentId")?.getFilterValue() as string) ??
                ""
              }
              onChange={(event) =>
                table
                  .getColumn("shipmentId")
                  ?.setFilterValue(event.target.value)
              }
              className="max-w-sm"
            />
            <Select
              defaultValue="ALL"
              onValueChange={(value) => {
                if (value === "ALL") {
                  table.getColumn("Status")?.setFilterValue(undefined);
                } else {
                  table.getColumn("Status")?.setFilterValue(value);
                }
              }}
            >
              <SelectTrigger>
                <SelectValue placeholder="Filter by Status" />
              </SelectTrigger>
              <SelectContent>
                <SelectItem value="ALL">All</SelectItem>
                <SelectItem value="DELIVERED">Delivered</SelectItem>
                <SelectItem value="CANCELLED">Cancelled</SelectItem>
                <SelectItem value="INCOMING_SHIPMENT">
                  Incoming Shipment
                </SelectItem>
                <SelectItem value="WAREHOUSE_RECEIVED">
                  Warehouse Received
                </SelectItem>
                <SelectItem value="WAREHOUSE_DEPARTED">
                  Warehouse Departed
                </SelectItem>
              </SelectContent>
            </Select>
          </div>
          <div className="flex items-center gap-2">
            <p className="text-sm text-muted-foreground">
              Last Refresh: {lastRefresh.toLocaleString()}
            </p>
            <Button
              variant="outline"
              size="sm"
              onClick={handleRefresh}
              disabled={isRefreshing}
            >
              <RotateCw
                className={cn("h-4 w-4", isRefreshing && "animate-spin")}
              />
            </Button>
          </div>
        </div>
        <Table>
          <TableHeader>
            {table.getHeaderGroups().map((headerGroup) => (
              <TableRow key={headerGroup.id}>
                {headerGroup.headers.map((header) => (
                  <TableHead key={header.id}>
                    {header.isPlaceholder
                      ? null
                      : flexRender(
                          header.column.columnDef.header,
                          header.getContext()
                        )}
                  </TableHead>
                ))}
              </TableRow>
            ))}
          </TableHeader>
          <TableBody>
            {table.getRowModel().rows?.length ? (
              table.getRowModel().rows.map((row) => (
                <TableRow key={row.id}>
                  {row.getVisibleCells().map((cell) => (
                    <TableCell key={cell.id}>
                      {flexRender(
                        cell.column.columnDef.cell,
                        cell.getContext()
                      )}
                    </TableCell>
                  ))}
                </TableRow>
              ))
            ) : (
              <TableRow>
                <TableCell
                  colSpan={columns.length}
                  className="h-24 text-center"
                >
                  No results.
                </TableCell>
              </TableRow>
            )}
          </TableBody>
        </Table>
        <div className="flex items-center justify-end space-x-2 py-4">
          <Button
            variant="outline"
            size="sm"
            onClick={() => table.previousPage()}
            disabled={!table.getCanPreviousPage()}
          >
            Previous
          </Button>
          <Button
            variant="outline"
            size="sm"
            onClick={() => table.nextPage()}
            disabled={!table.getCanNextPage()}
          >
            Next
          </Button>
        </div>
      </div>
    </div>
  );
}
