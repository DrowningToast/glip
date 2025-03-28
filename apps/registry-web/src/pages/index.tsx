import { useEffect } from "react";
import { useNavigate } from "react-router";

import { useQuery } from "@tanstack/react-query";
import { registryApi } from "../api/registry";

import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "../components/ui/table";
import WarehouseDialog from "../components/WarehouseDialog";
const Page = () => {
  const { data } = useQuery({
    queryKey: ["warehouses"],
    queryFn: async () => {
      const res = await registryApi.warehouseConnection.listConnections({
        headers: {
          authorization: `${import.meta.env.VITE_REGISTRY_SECRET_KEY || ""}`,
          authtype: "ADMIN",
        },
      });

      switch (res.status) {
        case 200:
          return res.body;
        case 400:
          throw new Error(res.body.message);
        case 500:
          throw new Error(res.body.message);
        default:
          throw new Error("Something went wrong");
      }
    },
  });

  const navigate = useNavigate();
  useEffect(() => {
    const isAuthenticated = localStorage.getItem("is_authenticated");
    if (!isAuthenticated) {
      navigate("/login");
    }
  }, [navigate]);

  return (
    <div className="w-full max-w-7xl mx-auto py-10">
      <div className="flex justify-between items-center">
        <h1 className="text-2xl font-bold">Warehouses</h1>
        <WarehouseDialog />
      </div>
      <Table className="mt-10">
        <TableHeader className="border-b border-gray-200">
          <TableHead>ID</TableHead>
          <TableHead>Warehouse ID</TableHead>
          <TableHead>Warehouse Name</TableHead>
          <TableHead>Status</TableHead>
          <TableHead>API Key</TableHead>
          <TableHead>Created At</TableHead>
          <TableHead>Updated At</TableHead>
          <TableHead>Last Used</TableHead>
        </TableHeader>
        <TableBody>
          {data?.result.items.map((warehouse) => (
            <TableRow
              key={warehouse.id}
              className="hover:bg-gray-50 border-b border-gray-200"
            >
              <TableCell>{warehouse.id}</TableCell>
              <TableCell>{warehouse.warehouse_id}</TableCell>
              <TableCell>{warehouse.name}</TableCell>
              <TableCell>{warehouse.status}</TableCell>
              <TableCell>{warehouse.api_key}</TableCell>
              <TableCell>
                {warehouse.created_at
                  ? new Date(warehouse.created_at).toLocaleString("th-TH", {
                      year: "numeric",
                      month: "2-digit",
                      day: "2-digit",
                      hour: "2-digit",
                      minute: "2-digit",
                    })
                  : "N/A"}
              </TableCell>
              <TableCell>
                {warehouse.updated_at
                  ? new Date(warehouse.updated_at).toLocaleString("th-TH", {
                      year: "numeric",
                      month: "2-digit",
                      day: "2-digit",
                      hour: "2-digit",
                      minute: "2-digit",
                    })
                  : "N/A"}
              </TableCell>
              <TableCell>
                {warehouse.last_used_at
                  ? new Date(warehouse.last_used_at).toLocaleString("th-TH", {
                      year: "numeric",
                      month: "2-digit",
                      day: "2-digit",
                      hour: "2-digit",
                      minute: "2-digit",
                    })
                  : "N/A"}
              </TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </div>
  );
};

export default Page;
