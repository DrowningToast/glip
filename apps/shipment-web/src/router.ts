// Generouted, changes to this file will be overridden
/* eslint-disable */

import { components, hooks, utils } from '@generouted/react-router/client'

export type Path =
  | `/`
  | `/create-tracking-form`
  | `/customer`
  | `/customer/create`
  | `/customer/create/form`
  | `/customer/shipment-list`
  | `/login`
  | `/login/admin`
  | `/login/admin/form`
  | `/login/customer`
  | `/login/customer/form`
  | `/login/role-selector`
  | `/login/warehouse`
  | `/login/warehouse/form`
  | `/track`
  | `/track/:id`
  | `/track/:id/result`

export type Params = {
  '/track/:id': { id: string }
  '/track/:id/result': { id: string }
}

export type ModalPath = never

export const { Link, Navigate } = components<Path, Params>()
export const { useModals, useNavigate, useParams } = hooks<Path, Params, ModalPath>()
export const { redirect } = utils<Path, Params>()
