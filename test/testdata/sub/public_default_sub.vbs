Public Default Sub Min(nOne, nTwo, Three)
    If nOne <= nTwo Then ' +1
        Min = nOne
        If nThree <= nOne Then ' +2
            Min  = nThree
        Else ' +1
            Min = nOne
        End If
    End If
End Sub