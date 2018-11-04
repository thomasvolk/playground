
(defun pi_step (cnt val num e) 
  (if (not (= 0 cnt))
    (pi_step (- cnt 1) 
        (+ val (/ num 
          (- (expt e 3) e)
        ))
        (* num -1)
        (+ e 2)
    )
    val
  )
)

(setq p (pi_step 200 3 4 3))
(format t "~% pi=~d " (float p))