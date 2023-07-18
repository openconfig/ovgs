# Ownership Voucher gRPC Service (ovgs)

This document describes the interface to used to access the Ownership Voucher
(OV) gRPC Server. The OV gRPC server allows administrators to manage users and
devices within an organization and enables users to request [RFC8366 ownership
vouchers](https://datatracker.ietf.org/doc/html/rfc8366) for devices they have
been assigned. The system is designed around gRPC calls.

The purpose of this service is to allow users to request ownership vouchers for
network element components (as identified by their serial number + ien) owned by
their organization. Organizations are divided into subgroups, and serial numbers
for products and users are assigned to groups within the organization. The
sections on authentication and authorization provide more details on this.

## Service Authentication

### Authentication

It is assumed that the vendor providing this service endpoint has a mechanism to
provide authentication for the service endpoint.  The following examples assume
the use of an OIDC-based authentication environment.

```mermaid
flowchart TB

    orgBlock["Org: AcmeCo
    Org ID: org-acmeco
    Users = [{username = admin, role = admin}]"]

    topGroup("Group: Default
    Group ID: group-e51b5c2c-eda1-4c27-8a86-ece7faab33f
    Domain Certs = [&lt;x509 domain cert 1&gt;, &lt;x509 domain cert 2&gt;]
    Components = [{ien = 300YY, serial_number = JGEXXXXXX}]
    Users = [{username = useracm, role = requestor}]")

    siteA("Group: SiteA
    Group ID: group-ad636092-fbc4-446c-ab2f-9df16719613a
    Components = [{ien = 300YY, serial_number = GACXXXXXX}]
    Users: [{username = userconsulting, role = admin}]")

    siteB("
    Group: SiteB
    Group ID: group-79ef6ff1-98d1-4fa7-98f3-7d98af4ee6ee
    Users: [{username = siteb, role = requestor}]")

    orgBlock --> topGroup
    topGroup --> siteA
    topGroup --> siteB
```

### Authorization

The above figure shows the relationships between different entities in the
system. This is the organization tree. The root of the tree is the organization
itself. In this example, the organization has one child group, which has two
child groups. Each group has a set of pinned domain certs, a set of users, and a
set of serials assigned to it. This system has 3 roles: **ADMIN**, **ASSIGNER**,
and **REQUESTOR**. An **ADMIN** can do everything that an **ASSIGNER** can do.
An **ASSIGNER** can do everything a **REQUESTOR** can do.

Each user has access to the group that it belongs to along with all of its
children, i.e., the user has access to the subtree rooted at the group that the
user belongs to. One user can belong to multiple groups and access multiple
subtrees in the hierarchy. The user has the same permission across the subtree
that it has been assigned at the root group of the subtree. If two subtrees
overlap for a user, the permission that gives the user higher access is applied.

In the example shown in the figure, the user with `username=siteb` cannot request
an ownership voucher for the serial JGEXXXXXX because the serial belongs to the
parent group. The user with username=useracm can request an ownership voucher
for the serial with serial number GACXXXXXX because the user has access to the
parent group of the serial.

## Serial Numbers

Ownership vouchers can be issued for components as identified by their
serial number + ien (hereafter simply referred to as serial number, the
inclusion of an ien as appropriate, is implied). Some products may also have a
TPM Public Key associated with them. By default, it is assumed that all the
customer-owned component serial numbers are added to the root group. Users with
the **ASSIGNER** or **ADMIN** role can move/add serials within the tree from
there. A serial can be assigned to a single group apart from the root group. See
the `/GetSerial` RPC for an example of the details available for a given serial
number.

## Pinned Domain Certificates (PDCs)

Pinned Domain Certificates (aka PDCs) are the roots of trust used in the
ownership voucher, along with "revocation checks" and an "expiry time". A
certificate ID is returned when a pinned domain cert is added to a group. While
creating ownership vouchers, this cert ID is used to reference the certificate.
See the `/GetDomainCert` RPC for an example of this data.

## Roles

Users are assigned to **ADMIN** or **REQUSTER** or **ASSIGNER** roles. An
**ADMIN** can do everything that a (**ASSIGNER**, **REQUESTOR**) can do, along
with requesting Ownership Vouchers. See `/GetGroup` RPC for an example. Role
assignment has relevance within the context of a group and associated group
hierarchy.

A vendor may also maintain specific internal roles for performing
support-related operations. An **ADMIN** role can set up the group hierarchy and
add users with different roles for different subtrees. An **ASSIGNER** can
create PDCs and associate them with groups, move serials between groups, and
issue ownership vouchers. A **REQUESTOR** role allows read access and the
ability to issue ownership vouchers.

<table>
  <tr>
   <td><strong>Role</strong>
   </td>
   <td><strong>Allowed RPCs</strong>
   </td>
   <td><strong>Allowed to be Assigned / Removed by</strong>
   </td>
   <td><strong>Allowed to be Assigned on</strong>
   </td>
   <td><strong>Allowed to be Viewed by</strong>
   </td>
  </tr>
  <tr>
   <td>SUPPORT
   </td>
   <td>All RPCs
   </td>
   <td>SUPPORT on the group/org or its parent
   </td>
   <td>the group/org or its children
   </td>
   <td>Anyone
   </td>
  </tr>
  <tr>
   <td>ADMIN
   </td>
   <td>All RPCs
   </td>
   <td>SUPPORT / ADMIN on the group/org or its parent
   </td>
   <td>the group/org or its children
   </td>
   <td>Anyone
   </td>
  </tr>
  <tr>
   <td>ASSIGNER
   </td>
   <td>All except CreateGroup, DeleteGroup, AddUserRole, RemoveUserRole,
   </td>
   <td>SUPPORT / ADMIN on the group/org or its parent
   </td>
   <td>the group/org or its children
   </td>
   <td>Anyone
   </td>
  </tr>
  <tr>
   <td>REQUESTOR
   </td>
   <td>GetGroup, GetUserRole, GetDomainCert, /≥≤, GetOwnershipVoucher
   </td>
   <td>SUPPORT / ADMIN / ASSIGNER on the group/org or its parent
   </td>
   <td>the group/org or its children
   </td>
   <td>Anyone
   </td>
  </tr>
</table>

## Groups

Groups are used to limit the scope of available data and allow for the members
of the group to manage a subset of items. Groups contain serials, domain certs,
users to roles assignments, and other groups as children. When an operator
begins using the ownership vouchers service, they will have a root group
created. Users can be assigned to more than one group and will have permission
to access their assigned and child groups. Child groups cannot access
information in their parent groups. A child can only have one parent in the
group hierarchy but can have multiple children. See the `/GetGroup` RPC for an
example of this data.

## Users

A user is uniquely identified using a combination of ID (username), type, and
the organization ID. The type of user supported today are either **USER** for a
SSO user, or **SERVICE_ACCOUNT** for a service account.

USER represents a user account, tied to their identity, as specified by the
authentication mechanism provided. SERVICE_ACCOUNT provides an identity that is
not tied to a human/user and is intended for programmatic access from any
app/service (say the bootstrap server) to the ovgs service. This can typically
be achieved by creating service accounts (with some name), and issuing long
lived tokens against that service account name.

## Bootstrapping an Organization

The process of bootstrapping an organization is expected to be vendor dependent
and is outside the scope of this service and document.

Note, it is assumed that the root group as defined by the vendor will have all
of operator's serial numbers assigned to it by default. Users then can modify
the assignment of the serials to different groups as needed.

## GRPC API

Users can use the following endpoints to manipulate the devices, groups, users,
roles, and permissions. Further details on each method can be found within the
protobuf.

The examples below assume that the org is **AcmeCo** with Org ID **org-acmeco**.

### /CreateGroup

- **Endpoint:** `/CreateGroup`
- **Minimum Role Needed:** admin
- **Endpoint Action:** Creates a named group as a child of an existing group or
  an organization.

#### Example Request

```text
parent = org-acmeco
description = acmeco-bu
```

#### Example Response

```text
group_id = group-e51b5c2c-eda1-4c27-8a86-ece7faa0dac2
```

### /GetGroup

- **Endpoint:** `/GetGroup`
- **Minimum Role Needed:** requestor
- **Endpoint Action:** For a named group, view the child groups, domain-certs
  IDs, serials and user/role mappings for that group.

#### Example Request

```text
group_id = group-e51b5c2c-eda1-4c27-8a86-ece7faa0dac2
```

#### Example Response

```text
group_id = group-e51b5c2c-eda1-4c27-8a86-ece7faa0dac2
cert_ids = [cert-7ccce4fc-1b28-469a-b4f5-79a4115d772b, ]
components = [{ien = 300XX, serial_number = JPEXXXX076}]
users = [{username = useracm, user_type = USER, org_id = org-acmeco, user_role = USER_ROLE_ADMIN}]
```

### /DeleteGroup

- **Endpoint:** `/DeleteGroup`
- **Minimum Role Needed:** admin
- **Endpoint Action:** Deletes a named group. Will refuse to do so if there are
  subgroups.

#### Example Request

```text
group_id = group-e51b5c2c-eda1-4c27-8a86-ece7faa0dac2
```

### /AddUserRole

- **Endpoint:** `/AddUserRole`
- **Minimum Role Needed:** admin
- **Endpoint Action:** Assigns a user to a role in a named group.

#### Example Request

```text
username = useracm
user_type = USER
org_id = org-acmeco
group_id = group-e51b5c2c-eda1-4c27-8a86-ece7faa0dac2
user_role = USER_ROLE_REQUESTOR
```

### /RemoveUserRole

- **Endpoint:** `/RemoveUserRole`
- **Minimum Role Needed:** admin
- **Endpoint Action:** Removes the role of the user in a named group. This
  essentially revokes a user’s access to the group.

#### Example Request

```text
username = useracm
user_type = USER
org_id = org-acmeco
group_id = group-e51b5c2c-eda1-4c27-8a86-ece7faa0dac2
```

### /GetUserRole

- **Endpoint:** `/GetUserRole`
- **Minimum Role Needed:** requestor
- **Endpoint Action:** For a given user, returns the permissions of the user
  across groups that it belongs to. A user can only view roles of another user
  in the groups that it has a role assigned to.

#### Example Request

```text
username = useracm
user_type = user
org_id = org-acmeco
```

#### Example Response

```text
groups = [{ group-e51b5c2c-eda1-4c27-8a86-ece7faa0dac2 = ADMIN}]
```

### /CreateDomainCert

- **Endpoint:** `/CreateDomainCert`
- **Minimum Role Needed:** assigner
- **Endpoint Action:** For a named group, set the values of:
  - Pinned-domain-cert
  - Domain-cert-revocation-checks
  - Expiry Time

#### Example Request

```text
group_id = group-e51b5c2c-eda1-4c27-8a86-ece7faa0dac2
certificate_der = <x509 cert ASN.1 der encoded>
revocation_checks = true
expiry_time = 2023-02-25T00:00:00.000Z
```

#### Example Response

```text
cert_id = cert-7ccce4fc-1b28-469a-b4f5-79a4115d772b
```

### /GetDomainCert

- **Endpoint:** `/GetDomainCert`
- **Minimum Role Needed:** requestor
- **Endpoint Action:** For a given domain cert id reveals the values of:
  - `group_id` that the cert belongs to
  - Pinned-domain-cert
  - Domain-cert-revocation-checks
  - `expiry_time`

#### Example Request

```text
cert_id = cert-7ccce4fc-1b28-469a-b4f5-79a4115d772b
```

#### Example Response

```text
group_id = group-e51b5c2c-eda1-4c27-8a86-ece7faa0dac2
certificate_der = <x509 ASN.1 der encoded>
revocation_checks = true
expiry_time = 2023-02-25T00:00:00.000Z
```

### /DeleteDomainCert

- **Endpoint:** `/DeleteDomainCert`
- **Minimum Role Needed:** assigner
- **Endpoint Action:** For a given cert_id, delete the cert from the database.

#### Example Request

```text
cert_id = cert-7ccce4fc-1b28-469a-b4f5-79a4115d772b
```

### /AddSerial

- **Endpoint:** `/AddSerial`
- **Minimum Role Needed:** assigner
- **Endpoint Action:** For a serial and named group, assign the serial to that group.

#### Example Request

```text
component = {ien = 300YY, serial_number = JPEXXXX1076}
group_id = group-e51b5c2c-eda1-4c27-8a86-ece7faa0dac2
```

### /RemoveSerial

- **Endpoint:** `/RemoveSerial`
- **Minimum Role Needed:** assigner
- **Endpoint Action:** For a serial and named group, remove the serial from the group.

#### Example Request

```text
component = {ien = 300YY, serial_number = JPEXXXX1076}
group_id = group-e51b5c2c-eda1-4c27-8a86-ece7faa0dac2
```

### /GetSerial

- **Endpoint:** `/GetSerial`
- **Minimum Role Needed:** requestor
- **Endpoint Action:** Given a serial number, return all the facts about the serial.

#### Example Request

```text
component = {ien = 300YY, serial_number = JPEXXXX1076}
```

#### Example Response

```text
public_key_der = <TPM public key, ASN.1 der encoded (if applicable for this part)>
group_ids = [group-e51b5c2c-eda1-4c27-8a86-ece7faa0dac2]
mac_addr = 00:00:5e:00:53:af
model = DCS-7800-SUP
```

### /GetOwnershipVoucher

- **Endpoint:** `/GetOwnershipVoucher`
- **Minimum Role Needed:** requestor
- **Endpoint Action:** Given a Serial Number, Domain Cert ID, IEN (IANA
  Enterprise Number of the device vendor, e.g., 30065 is Arista’s IEN) and OV
  lifetime this endpoint will do the following:

  - Verify that the requestor has access to the device with the serial number.
  - Verify that the requested lifetime in not in the past and is within the cert
    expiry time
  - Issue an OV with the requested and set parameters.
  - Also return the TPM public key of the device for downstream validation or
    other purposes if available.
  - The TPM public key is not required to be used, but allows the customer to
    perform additional in-depth validations of the product they receive.

#### Example Request

```text
component = {ien = 300YY, serial_number = JPE29451076}
cert_id = cert-7ccce4fc-1b28-469a-b4f5-79a4115d772b
lifetime = 2023-02-25T00:00:00.000Z
```

#### Example Response

```text
voucher_cms = <voucher, example below>
public_key_der = <switch TPM public key, ASN.1 der encoded, example below>
```

## Service Use Examples

The following examples assume the organization has been bootstrapped by the
vendor and an `admin` user has been created.

Once the admin user has logged in and specified their username, they shall be
added with the role **ADMIN** to the root of the organization tree.

It is recommended that the admin user create a service account, give it an ADMIN
role over the organization and tree, and then use the service account for
interacting with the ownership voucher service programmatically, including
setting up their organization tree. The organization acmeco will be used for the
examples below.

The following tree with just one node is set up when a baseline organization or
tenant is created for an operator. For example, this node is set up for the
`acmeco organization` with an admin user.

```mermaid
%%{init: {"flowchart": {"htmlLabels": false}} }%%
flowchart TB

    orgBlock("org_id: org-acmeco
users = [{username = admin, user_type = user, user_role = admin}]
switch_serials = [ABC101, ABC102]
")

orgBlock
```

### Setup

#### Install grpcurl

`grpcurl` can be either compiled from the source code available on  Github or
binaries can be directly downloaded from
<https://github.com/fullstorydev/grpcurl/releases>.

Subsequent examples use `grpcurl`, where the fields are encoded in JSON.

### Getting the Protobuf File

The protobuf file can be found in the following github repository:
<https://github.com/openconfig/ovgs/>

## User Workflow

The operator will need to use both their vendor-specific UI and the `grpcurl`
tool (or something equivalent) to interact with the Ownership Voucher Grpc
service and perform necessary operations to get access to the vouchers, as
explained below.

## Getting the Token

Obtaining an API access token will be a function of the mechanisms provided by the
respective vendor.  Review your vendor's documentation for the relevant details.

## Creating A New User

To create a new user or service account, follow the steps provided by the vendor
for interaction with their respective account management processes.

### Creating a Group

Create a group with the name "default" as a child to the root group; as
shown in the picture below. Copy the group ID from the response, this will be
required later.

```mermaid
%%{init: {"flowchart": {"htmlLabels": false}} }%%
flowchart TB

    orgBlock("org_id: org-acmeco
users = [{username = admin, user_type = user, user_role = admin}]
switch_serials = [ABC101, ABC102]
")
    defaultGroup("Group Name: default
group_id: group-3e7e2431-6c73-423b-91ef-b734a13daaab
")
orgBlock --> defaultGroup
```

```shell
$ ACCESS_TOKEN=<token of admin>
$ grpcurl -H "Cookie: access_token=${ACCESS_TOKEN}"         \
    -proto ovgs.proto                                       \
    -d '{"parent": "org-acmeco", "description": "default"}' \
    www.a-network-vendor.io:443                             \
    ovgs.v1.OwnershipVoucherService/CreateGroup
```

#### Response

```text
{"group_id": "group-3e7e2431-6c73-423b-91ef-b734a13daaab"}
```

### Adding New Service Account to Organization Tree

To add the `srv-admin` service account to the organization tree add the service
account to the root of the organization tree. Notice that the following is set:

1. The username is set to `srv-admin`
2. `user_type` as `ACCOUNT_TYPE_SERVICE_ACCOUNT` (`ACCOUNT_TYPE_USER` for user
   accounts, `ACCOUNT_TYPE_SERVICE_ACCOUNT` for service accounts)
3. `org_id` is `org-acmeco`
4. `group_id` is same as `org_id` because the account is assigned to the root group
5. user_role is set to `USER_ROLE_ADMIN`

```shell
$ ACCESS_TOKEN=<token of admin>
$ grpcurl -H "Cookie: access_token=${ACCESS_TOKEN}" \
    -proto ovgs.proto                               \
    -d '{"username": "srv-admin", "user_type": "ACCOUNT_TYPE_SERVICE_ACCOUNT", "org_id": "org-acmeco", "group_id": "org-acmeco", "user_role": "USER_ROLE_ADMIN"}' \
  www.a-network-vendor.io:443                       \
  ovgs.v1.OwnershipVoucherService/AddUserRole
```

#### Response

```text
{}
```

### Adding New User Account to Organization Tree

Once this command is executed successfully, the user is added to the root group
and this token can be used for running `grpcurl` commands. Adding New User Account
to Organization Tree To add the user-admin-on-default user account to our
organization tree. Add this user account to the group named "default". Notice
that the following are set:

1. The username to `user-admin-on-default`
2. user_type as `ACCOUNT_TYPE_USER` (`ACCOUNT_TYPE_USER` for user accounts,
   `ACCOUNT_TYPE_SERVICE_ACCOUNT` for service accounts)
3. `org_id` to `org-acmeco`
4. `group_id` to `group-3e7e2431-6c73-423b-91ef-b734a13daaab`, that is the group
   ID of the "default" group
5. user_role to set to USER_ROLE_ADMIN

```shell
$ ACCESS_TOKEN=<token of srv-admin>
$ grpcurl -H "Cookie: access_token=${ACCESS_TOKEN}" \
    -proto ovgs.proto                               \
    -d '{"username": "user-admin-on-default", "user_type": "ACCOUNT_TYPE_USER", "org_id": "org-acmeco", "group_id": "group-3e7e2431-6c73-423b-91ef-b734a13daaab", "user_role": "ADMIN"}' \
    www.a-network-vendor.io:443                     \
    ovgs.v1.OwnershipVoucherService/AddUserRole
```

#### Response

```text
{}
```

Once this command is executed successfully, the user is added to the "default"
group.

Let’s also create a service account srv-admin-on-default and add it to the
"default" group. The organization tree is now structured as follows:

```mermaid
%%{init: {"flowchart": {"htmlLabels": false}} }%%
flowchart TB

    orgBlock("
    org_id: org-acmeco
    users = [
      {username: admin, user_type: user, user_role: admin},
      {username: srv-admin, user_type: account, user_role: admin}
    ]
    switch_serials = [ABC101, ABC102]
    ")

    defaultGroup("
    Group Name: default
    group_id: group-3e7e2431-6c73-423b-91ef-b734a13daaab
    users = [
        {username: user-admin-on-default, user_type: user, user_role: admin},
        {username: srv-admin-on-default, user_type: account, user_role: admin}
      ]
    ")

    orgBlock --> defaultGroup
```

### Viewing Roles of A User

Let’s try to view the roles of the srv-admin user. Here when using the access
token of the srv-admin user, the following response is seen:

```shell
$ ACCESS_TOKEN=<token of srv-admin>
$ grpcurl -H "Cookie: access_token=${ACCESS_TOKEN}" \
    -proto ovgs.proto                               \
    -d '{"username": "srv-admin", "user_type": "ACCOUNT_TYPE_SERVICE_ACCOUNT", "org_id": "org-acmeco"}' \
    www.a-network-vendor.io:443                     \
    ovgs.v1.OwnershipVoucherService/GetUserRole
```

#### Response

```text
{"groups": {"org-acmeco": "ADMIN"}}
```

Now, if use access token of srv-admin-on-default is used, no roles will be seen.

```shell
$ ACCESS_TOKEN=<token of srv-admin-on-default>
$ grpcurl -H "Cookie: access_token=${ACCESS_TOKEN}" \
    -proto ovgs.proto                               \
   -d '{"username": "srv-admin", "user_type": "ACCOUNT_TYPE_SERVICE_ACCOUNT", "org_id": "org-acmeco"}' \
   www.a-network-vendor.io:443                      \
 ovgs.v1.OwnershipVoucherService/GetUserRole
```

#### Response

```text
{}
```

### Adding A Pinned Domain Certificate

```shell
 ACCESS_TOKEN=<token of srv-admin-on-default>
$ grpcurl -H "Cookie: access_token=${ACCESS_TOKEN}" \
    -proto ovgs.proto                               \
    -d '{"group_id": "group-3e7e2431-6c73-423b-91ef-b734a13daaab", "certificate_der": "MIIHRzCCss { ... snipped ... } OEGsiDoRSlA==", "revocation_checks": true, "expiry_time": "2023-02-25T00:00:00.000Z"}' \
    www.a-network-vendor.io:443                     \
    ovgs.v1.OwnershipVoucherService/CreateDomainCert
```

#### Response

```text
{"cert_id": "cert-29466354-a669-4c47-91cf-f214c03626db"}
```

### Retrieving a Pinned Domain Cert

```shell
$ ACCESS_TOKEN=<token of srv-admin-on-default>
$ grpcurl -H "Cookie: access_token=${ACCESS_TOKEN}"             \
  -proto ovgs.proto                                             \
  -d '{"cert_id": "cert-29466354-a669-4c47-91cf-f214c03626db"}' \
  www.a-network-vendor.io:443                                   \
  ovgs.v1.OwnershipVoucherService/CreateDomainCert
```

#### Response

```text
{
  "group_id": "group-3e7e2431-6c73-423b-91ef-b734a13daaab",
  "certificate_der": "MIIHRzCCBi+gAwIB...I6dj87VD+laMUBd7HBtOEGsiDoRSlA==",
  "revocation_checks": true,
  "expiry_time": "2023-02-25T00:00:00.000Z"
}
```

### Getting All Serial Numbers

The operator can use the `/GetGroup` RPC with `group_id` set to the `org_id`. For
example, if `acmeco` wants to get all of their serial numbers that they own and
their `org_id` is `org-acmeco`, they can run:

```shell
$ ACCESS_TOKEN=<token of srv-admin>
$ grpcurl -H "Cookie: access_token=${ACCESS_TOKEN}" \
    -proto ovgs.proto                               \
    -d '{"group_id": "org-acmeco"}'                 \
    www.a-network-vendor.io:443                     \
    ovgs.v1.OwnershipVoucherService/GetGroup
```

#### Response

```javascript
{
  "components": [{
          "ien": "30065",
          "serialNumber": "ABC101"
    }, {
          "ien": "30065",
          "serialNumber": "ABC102"
    }],
  "users": [{
          "username": "admin",
          "user_type": "ACCOUNT_TYPE_USER",
          "org_id": "org-acmeco",
          "user_role": "USER_ROLE_ADMIN"
    }, {
          "username": "srv-admin",
          "user_type": "ACCOUNT_TYPE_SERVICE_ACCOUNT",
          "org_id": "org-acmeco",
          "user_role": "USER_ROLE_ADMIN"
    }],
  "child_group_ids": [ "group-3e7e2431-6c73-423b-91ef-b734a13daaab"]
  ...
}
```

### Assigning A Serial Number to A Group

```shell
$ ACCESS_TOKEN=<token of srv-admin>
$ grpcurl -H "Cookie: access_token=${ACCESS_TOKEN}"               \
    -proto ovgs.proto                                             \
    -d '{"component": {"ien": "30065","serial_number": "ABC101"}, \
    "group_id": "group-3e7e2431-6c73-423b-91ef-b734a13daaab"}'    \
    www.a-network-vendor.io:443                                   \
    ovgs.v1.OwnershipVoucherService/AddSerial
```

#### Response

```text
{}
```

### Getting Public Key of the TPM for A Serial Number (if applicable)

```shell
$ ACCESS_TOKEN=<token of srv-admin>
$ grpcurl -H "Cookie: access_token=${ACCESS_TOKEN}" 		   \
    -proto ovgs.proto                               		   \
    -d '{"component": {"ien": "30065","serial_number": "ABC101"}}' \
    www.a-network-vendor.io:443                     		   \
    ovgs.v1.OwnershipVoucherService/GetSerial
```

#### Response

```text
{
  "public_key_der": "MIIBIjANBgkqhkiG9w0BA...A4M3QIDAQAB",
  "group_ids": [ "group-3e7e2431-6c73-423b-91ef-b734a13daaab", "org-acmeco" ],
  "mac_addr": "00:00:5e:00:53:af",
  "model": "DCS-7800-SUP"
}
```

### Getting Details of A Group

```shell
$ ACCESS_TOKEN=<token of srv-admin-on-default>
$ grpcurl -H "Cookie: access_token=${ACCESS_TOKEN}"                 \
    -proto ovgs.proto                                               \
    -d '{"group_id": "group-3e7e2431-6c73-423b-91ef-b734a13daaab"}' \
    www.a-network-vendor.io:443                                     \
    ovgs.v1.OwnershipVoucherService/GetGroup

```

#### Response

```text
{
  "cert_ids": [ "cert-29466354-a669-4c47-91cf-f214c03626db" ],
  "serial_numbers": [ "ABC101" ],
  "users": [{
          "username": "srv-admin-on-default",
          "user_type": "ACCOUNT_TYPE_SERVICE_ACCOUNT",
          "org_id": "org-acmeco",
          "user_role": "USER_ROLE_ADMIN"
    }, {
          "username": "user-admin-on-default",
          "user_type": "ACCOUNT_TYPE_SERVICE_ACCOUNT",
          "org_id": "org-acmeco",
          "user_role": "USER_ROLE_ADMIN"
    }]
   ...
}
```

### Getting Ownership Voucher

```shell
$ ACCESS_TOKEN=<token of srv-admin-on-default>
$ grpcurl -H "Cookie: access_token=${ACCESS_TOKEN}"               \
    -proto ovgs.proto                                             \
    -d '{"component": {"ien": "30065","serial_number": "ABC101"}, "cert_id": "cert-29466354-a669-4c47-91cf-f214c03626db", "lifetime": "2023-02-25T00:00:00.000Z"}' \
    www.a-network-vendor.io:443                                   \
    ovgs.v1.OwnershipVoucherService/GetOwnershipVoucher
```

#### Response

```text
{ "voucher_cms": "MIIeYAYJKoZI...XyC/dx8DbRGBWKK/pcGG+U50PRt86Q==",
    "public_key_der": "MIIBIjANBg...BVdgA4M3QIDAQAB" }
```

Note that every time a voucher is requested for a device with the same PDC and
lifetime, a different voucher is generated. All of the vouchers are valid
vouchers.

### Remove a User from a Group

```shell
$ ACCESS_TOKEN=<token of srv-admin>
$ grpcurl -H "Cookie: access_token=${ACCESS_TOKEN}"               \
    -proto ovgs.proto                                             \
    -d '{"username": "user-admin-on-default",                     \
       "user_type": "ACCOUNT_TYPE_USER",                          \
       "org_id": "org-acmeco",                                    \
       "group_id": "group-3e7e2431-6c73-423b-91ef-b734a13daaab"}' \
    www.a-network-vendor.io:443                                   \
    ovgs.v1.OwnershipVoucherService/RemoveUserRole
```

#### Response

```text
{}
```

### Deleting A User or A Service Account

Deletion of a user or a service account is handled through the mechanisms
provided by the vendor's user management interface.

### Deleting A Serial Number from A Group

```shell
$ ACCESS_TOKEN=<token of srv-admin-on-default>
$ grpcurl -H "Cookie: access_token=${ACCESS_TOKEN}"               \
    -proto ovgs.proto                                             \
    -d '{"component": {"ien": "30065","serial_number": "ABC101"}, \
    "group_id": "group-3e7e2431-6c73-423b-91ef-b734a13daaab"}'    \
    www.a-network-vendor.io:443                                   \
    ovgs.v1.OwnershipVoucherService/RemoveSerial
```

#### Response

```text
{}
```

### Deleting A Pinned Domain Cert

```shell
$ ACCESS_TOKEN=<token of srv-admin-on-default>
$ grpcurl -H "Cookie: access_token=${ACCESS_TOKEN}"               \
    -proto ovgs.proto                                             \
    -d '{"cert_id": "cert-29466354-a669-4c47-91cf-f214c03626db"}' \
    www.a-network-vendor.io:443                                   \
    ovgs.v1.OwnershipVoucherService/DeleteDomainCert
```

#### Response

```text
{}
```

### Deleting A Group

Deleting a group is only allowed when there are no objects such as certificates,
device serials, children groups, and users associated with the group.

```shell
$ ACCESS_TOKEN=<token of admin>
$ grpcurl -H "Cookie: access_token=${ACCESS_TOKEN}"               \
-proto ovgs.proto                                                 \
  -d '{"group_id": "group-3e7e2431-6c73-423b-91ef-b734a13daaab"}' \
  www.a-network-vendor.io:443                                     \
  ovgs.v1.OwnershipVoucherService/DeleteGroup
```

#### Response

```text
{}
```

## Format of an Ownership Voucher

More details are available at [RFC8366
Section-5.3](https://tools.ietf.org/html/rfc8366#section-5.3).

- `created-on` and `expires-on` follow RFC3339 format
- `expires-on` is omitted if the certificate is not expected to expire.
- `pinned-domain-cert` is ASN.1 DER encoded x509 certificate converted into
  string using the base64 encoding
- `domain-cert-revocation-checks` is going to be a boolean value

```text
{
  "ietf-voucher:voucher": {
        "created-on": "2023302-11T13:45:31.69401473+05:30",
        "expires-on": "2023-08-11T13:45:31+05:30",
        "serial-number": "JPEXXXX27",
        "assertion": "verified",
        "pinned-domain-cert": "MIIFjjC {... snipped ...} HBUoCj0M6oIjhTcvHQ==",
        "domain-cert-revocation-checks": true
  }
}
```

The ownership voucher is signed and provided in binary CMS format as explained
in [rfc8366](https://tools.ietf.org/html/rfc8366) Note that the ownership
vouchers are not encrypted, only signed by the vendor.
