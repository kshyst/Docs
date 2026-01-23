```sql
create trigger Tashvighi
    after insert on call
    referencing newCall as newCall
    for each row
    begin
        declare int count;

        if newCall.duration > 30 then
            if exists (
                select 1
                from customer as c1
                join customer as c2
                where c1.customerid = newCall.callerid
                  and c1.planid = 100
                  and c2.customerid = newCall.calledid
                  and c2.planid = 100
            ) then
                select count(*)
                into count
                from call
                join customer on calledid = customer.customerid
                where callerid = newCall.callerid
                  and planid = 100
                  and duration > 30
                  and callid <> newCall.callid;

                if count / 5 < (count + 1) / 5 then
                    update runningbill
                    set amount = amount * 0.98
                    where customerid = newCall.callerid;
                end if;
            end if;
        end if;
    end;

```

```sql
create trigger UpdateRunningBill
    after insert on call
    referencing new as newCall
    for each row
    update runningbill
        set amount + (
            select  priceplan.connectionfee + 
                    (priceplan.pricepersecond * newCall.duration)
            from priceplan
            join customer on priceplan.planid = customer.planid
            where customer.customerid = newCall.callerId
            )
    where runningbill.customerid = newCall.callerID;
```

```sql
create assertion No200ForYou
    check(
       not exists(
            select from 1 from customer
            where customer.planid = 200 and
            360000 > (
                select sum(duration) from call
                group by callerid, calledid
                where (callerid = customer.customerid or 
                       calledid = customer.customerid) and
                timestamp > NOW() - INTERVAL '1 year'
            )
        )
    );
       
```