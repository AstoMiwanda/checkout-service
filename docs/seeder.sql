INSERT INTO public.customers (id, name, email, created_at, updated_at, deleted_at)
VALUES (DEFAULT, 'Asto', 'asto@mail.com', DEFAULT, null, null);

INSERT INTO public.products (id, sku, name, price, is_active, created_at, updated_at, deleted_at)
VALUES ('8e81c473-170d-4f1e-9f08-772cc1e318e8', '120P90', 'Google Home', 49.99, true, '2025-04-26 21:15:05.205913',
        '2025-04-26 21:15:05.208274', null);

INSERT INTO public.products (id, sku, name, price, is_active, created_at, updated_at, deleted_at)
VALUES ('b84b34b9-b616-4aab-97b1-bda15ca275a3', '43N23P', 'Macbook Pro', 5399.99, true, '2025-04-26 21:15:39.504781',
        '2025-04-26 21:15:39.505734', null);

INSERT INTO public.products (id, sku, name, price, is_active, created_at, updated_at, deleted_at)
VALUES ('558c4bcb-71f3-4256-ac5b-89eaa4748c38', 'A304SD', 'Alexa Speaker', 109.5, true, '2025-04-26 21:16:13.973891',
        '2025-04-26 21:16:13.974870', null);

INSERT INTO public.products (id, sku, name, price, is_active, created_at, updated_at, deleted_at)
VALUES ('14dc8976-ccea-43b1-9611-0816bf51a3a9', '234234', 'RaspBerry Pi B', 30, true, '2025-04-26 21:16:44.080474',
        '2025-04-26 21:16:44.081494', null);


INSERT INTO public.discounts (id, type, description, discount_value, is_active, valid_from, valid_to, created_at,
                              updated_at, deleted_at)
VALUES ('e4b0b7ce-9e52-4614-8451-8d844d90f2b1', 'buy_x_get_y', 'Buy a Macbook Pro get Raspberry Pi B', null, null, null,
        '2025-04-23 21:23:52.000000', '2025-04-28 21:24:03.000000', '2025-04-26 14:24:10.168585', null);

INSERT INTO public.discounts (id, type, description, discount_value, is_active, valid_from, valid_to, created_at,
                              updated_at, deleted_at)
VALUES ('5b8889c6-fe27-4cf7-8bcb-ab0ac8812c7b', 'buy_x_pay_y', 'Buy 3 Google Home Pay price of 2', null, null, null,
        '2025-04-23 21:23:52.000000', '2025-04-28 21:24:03.000000', '2025-04-26 14:24:10.168585', null);

INSERT INTO public.discounts (id, type, description, discount_value, is_active, valid_from, valid_to, created_at,
                              updated_at, deleted_at)
VALUES ('64b0c34a-abf4-4ba2-bf6a-c1f20d34e73b', 'bulk_discount', 'Buy more than 3 Alexa get a 10% discount', null, null,
        null, '2025-04-23 21:23:52.000000', '2025-04-28 21:24:03.000000', '2025-04-26 14:24:10.168585', null);

INSERT INTO public.discount_rules (id, discount_id, product_id, role, quantity, quantity_operator, created_at,
                                   updated_at, deleted_at)
VALUES ('560e0a1e-bc20-4ccd-ba8d-caf4915cf4bc', 'e4b0b7ce-9e52-4614-8451-8d844d90f2b1',
        'b84b34b9-b616-4aab-97b1-bda15ca275a3', 'buy', 1, 'equal', '2025-04-26 14:28:10.011701', null, null);

INSERT INTO public.discount_rules (id, discount_id, product_id, role, quantity, quantity_operator, created_at,
                                   updated_at, deleted_at)
VALUES ('feb5a0b4-4009-421f-8c5c-5f9009baee4c', 'e4b0b7ce-9e52-4614-8451-8d844d90f2b1',
        '14dc8976-ccea-43b1-9611-0816bf51a3a9', 'get', 1, 'equal', '2025-04-26 14:28:10.011701', null, null);

INSERT INTO public.discount_rules (id, discount_id, product_id, role, quantity, quantity_operator, created_at,
                                   updated_at, deleted_at)
VALUES ('b615177d-a368-458c-a560-00419a203ad0', '5b8889c6-fe27-4cf7-8bcb-ab0ac8812c7b',
        '8e81c473-170d-4f1e-9f08-772cc1e318e8', 'buy', 3, 'equal', '2025-04-26 14:29:22.166641', null, null);

INSERT INTO public.discount_rules (id, discount_id, product_id, role, quantity, quantity_operator, created_at,
                                   updated_at, deleted_at)
VALUES ('baa294f8-4007-4f04-9975-d26d8692a50e', '5b8889c6-fe27-4cf7-8bcb-ab0ac8812c7b',
        '8e81c473-170d-4f1e-9f08-772cc1e318e8', 'discount', 1, 'equal', '2025-04-26 14:29:22.166641', null, null);

INSERT INTO public.discount_rules (id, discount_id, product_id, role, quantity, quantity_operator, created_at,
                                   updated_at, deleted_at)
VALUES ('e210b023-626a-4263-a2eb-bb1a9f0c9525', '64b0c34a-abf4-4ba2-bf6a-c1f20d34e73b',
        '558c4bcb-71f3-4256-ac5b-89eaa4748c38', 'buy', 3, 'more_than', '2025-04-26 14:32:29.302512', null, null);

