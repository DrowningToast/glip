/*
  Warnings:

  - Changed the type of `shipmentId` on the `Shipments` table. No cast exists, the column would be dropped and recreated, which cannot be done if there is data, since the column is required.

*/
-- AlterTable
ALTER TABLE "Shipments" DROP COLUMN "shipmentId",
ADD COLUMN     "shipmentId" INTEGER NOT NULL;
