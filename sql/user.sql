create database `leet_watcher`;

use `leet_watcher`;

create table if not exists `lw_user` (
    `id` bigint(20) not null primary key auto_increment comment 'ID',
    `username` varchar(50) not null comment 'LeetCode用户名',
    `remark` varchar(50) not null default '' comment '备注',
    `last_submit_time` timestamp null comment '最后一次提交时间',
    `create_time` timestamp default current_timestamp() comment '创建时间',
    `update_time` timestamp default current_timestamp() on update current_timestamp() comment '更新时间'
);
