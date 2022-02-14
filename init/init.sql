create schema lts_goods collate utf8mb4_0900_ai_ci;
create schema lts_orders collate utf8mb4_0900_ai_ci;
create schema lts_users collate utf8mb4_0900_ai_ci;

-- auto-generated definition
create table lts_goods.goods
(
    id           int unsigned auto_increment
        primary key,
    goods_name   varchar(100)                           not null comment '商品名称',
    price        int unsigned                           not null comment '价格（分）',
    description  varchar(100) default ''                not null comment '商品描述',
    stock_count  int unsigned default '0'               not null comment '库存数量',
    frozen_count int unsigned default '0'               not null comment '冻结数量',
    mtime        datetime     default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '修改时间',
    ctime        datetime     default CURRENT_TIMESTAMP not null comment '创建时间'
)
    comment '商品信息表';

create table lts_goods.goods_detail
(
    id           int unsigned auto_increment primary key,
    good_id      int unsigned                           not null comment '商品id',
    stock_count  int unsigned default '0'               not null comment '库存数量',
    frozen_count int unsigned default '0'               not null comment '冻结数量',
    type         tinyint unsigned                       not null comment '',
    ctime        datetime     default CURRENT_TIMESTAMP not null comment '创建时间'
)
    comment '商品库存流水表';


create table lts_orders.orders
(
    id            int auto_increment
        primary key,
    state         tinyint unsigned                   not null comment '状态。1、try。2、commit。3、cancel。',
    order_no      varchar(100)                       not null comment '订单号，业务需要',
    user_id       int unsigned                       not null comment '用户id',
    goods_id      int unsigned                       not null comment '商品id',
    goods_count   int unsigned                       not null comment '购买商品数量',
    change_amount int                                not null comment '订单金额',
    mtime         datetime default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP,
    ctime         datetime default CURRENT_TIMESTAMP not null
)
    comment '订单表';

-- auto-generated definition
create table lts_orders.orders_detail
(
    id                 int unsigned auto_increment
        primary key,
    orders_id          int unsigned                       not null comment '订单id',
    type               tinyint unsigned                   not null comment '操作类型。1、try。2、commit。3、cancel。',
    user_id            int unsigned                       not null comment '用户id',
    user_balance       int unsigned                       not null comment '用户余额',
    user_frozen_amount int unsigned                       not null comment '用户冻结金额',
    mtime              datetime default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '修改时间',
    ctime              datetime default CURRENT_TIMESTAMP not null comment '创建时间'
) comment '订单操作表';


-- auto-generated definition
create table lts_users.users
(
    id            int unsigned auto_increment
        primary key,
    username      varchar(100)                           not null comment '用户名',
    balance       int unsigned default '0'               not null comment '余额，单位：分',
    frozen_amount int unsigned default '0'               null comment '冻结金额，单位分',
    mtime         datetime     default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '修改时间',
    ctime         datetime     default CURRENT_TIMESTAMP not null comment '创建时间'
)
    comment '用户表';

create table lts_users.users_detail
(
    id            int unsigned auto_increment primary key,
    user_id       int unsigned                       not null comment '用户id',
    balance       int unsigned                       not null comment '余额',
    frozen_amount int unsigned                       not null comment '冻结金额',
    type          tinyint unsigned                   not null comment '',
    ctime         datetime default CURRENT_TIMESTAMP not null comment '创建时间'
)
    comment '用户余额流水表';
