-- CreateEnum
CREATE TYPE "ShipmentStatus" AS ENUM ('INCOMING_SHIPMENT', 'WAREHOUSE_RECEIVED', 'WAREHOUSE_DEPARTED', 'DELIVERED', 'CANCELLED');

-- CreateTable
CREATE TABLE "User" (
    "id" TEXT NOT NULL,
    "email" TEXT NOT NULL,
    "password" TEXT NOT NULL,

    CONSTRAINT "User_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "Shipments" (
    "id" TEXT NOT NULL,
    "shipmentId" TEXT NOT NULL,
    "route" TEXT[],
    "last_warehouse_id" TEXT,
    "departure_warehouse_id" TEXT NOT NULL,
    "departure_address" TEXT,
    "destination_warehouse_id" TEXT NOT NULL,
    "destination_address" TEXT NOT NULL,
    "created_by" INTEGER NOT NULL,
    "owner_id" INTEGER NOT NULL,
    "status" "ShipmentStatus" NOT NULL,
    "total_weight" DOUBLE PRECISION NOT NULL,
    "total_volume" DOUBLE PRECISION NOT NULL,
    "special_instructions" TEXT,
    "created_at" TIMESTAMP(3) NOT NULL,
    "updated_at" TIMESTAMP(3) NOT NULL,
    "from_warehouse_id" TEXT NOT NULL,
    "to_warehouse_id" TEXT NOT NULL,
    "delivery_time" TIMESTAMP(3),

    CONSTRAINT "Shipments_pkey" PRIMARY KEY ("id")
);

-- CreateIndex
CREATE UNIQUE INDEX "User_email_key" ON "User"("email");
